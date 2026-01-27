// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.
// NOTE: PreFetch and throttling additions are manual - preserve on regeneration.

package requestconfig

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

// SDK_VERSION is logged on first OAuth2 token fetch to verify which code is deployed
// UPDATE THIS EVERY TIME YOU MODIFY THIS FILE
const SDK_VERSION = "2026-01-27-v16-prefetch-throttle"

var sdkVersionLogged sync.Once

var OAuth2Cache = newOAuth2Cache("/v2/auth/token")

func newOAuth2Cache(tokenUrls ...string) map[string]*OAuth2State {
	state := make(map[string]*OAuth2State, len(tokenUrls))
	for _, url := range tokenUrls {
		state[url] = &OAuth2State{authPath: url}
	}
	return state
}

// OAuth2State represents the OAuth2 provider configuration and manages token
// caching with support for pre-fetching and failure throttling.
type OAuth2State struct {
	authPath string
	entries  sync.Map // map[oAuth2ClientCredentials]*oAuthTokenInfo

	// Failure throttling to prevent retry storms
	failureMu    sync.Mutex
	lastFailedAt time.Time
}

// FailureCooldown is how long to wait after a token fetch failure before retrying.
// This prevents thundering herd when auth service is overwhelmed.
const FailureCooldown = 5 * time.Second

type oAuth2ClientCredentials struct {
	ClientID     string
	ClientSecret string
}

type oAuthTokenInfo struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`

	expiresAt     time.Time
	reqTokenMutex sync.Mutex
}

func (tok *oAuthTokenInfo) reset() {
	tok.AccessToken = ""
	tok.ExpiresIn = 0
	tok.expiresAt = time.Time{}
}

func (auth *oAuthTokenInfo) validToken() string {
	if auth.AccessToken != "" &&
		time.Now().Before(auth.expiresAt.Add(-10*time.Second)) {
		return auth.AccessToken
	}
	return ""
}

// PreFetch fetches a token proactively before parallel operations begin.
// This should be called once at startup (e.g., in Terraform provider Configure)
// to avoid thundering herd when many parallel operations all need tokens.
// It retries with exponential backoff on failure.
func (state *OAuth2State) PreFetch(
	ctx context.Context,
	baseURL *url.URL,
	clientID, clientSecret string,
	httpClient *http.Client,
) error {
	// Log SDK version
	sdkVersionLogged.Do(func() {
		fmt.Fprintf(
			os.Stderr,
			"[SDK] serval-go %s | PreFetch starting\n",
			SDK_VERSION,
		)
	})

	tokenInfo := state.findByCredentials(clientID, clientSecret)

	// Already have a valid token
	if tokenInfo.validToken() != "" {
		fmt.Fprintf(
			os.Stderr,
			"[SDK PreFetch] Token already cached and valid\n",
		)
		return nil
	}

	fmt.Fprintf(os.Stderr, "[SDK PreFetch] Fetching token with retries...\n")

	// Retry with exponential backoff
	var lastErr error
	for attempt := 0; attempt < 5; attempt++ {
		if attempt > 0 {
			backoff := time.Duration(
				1<<attempt,
			) * 100 * time.Millisecond // 200ms, 400ms, 800ms, 1.6s
			fmt.Fprintf(
				os.Stderr,
				"[SDK PreFetch] Retry attempt %d after %v\n",
				attempt+1,
				backoff,
			)
			time.Sleep(backoff)
		}

		err := state.fetchTokenDirect(
			ctx,
			baseURL,
			clientID,
			clientSecret,
			httpClient,
			tokenInfo,
		)
		if err == nil {
			fmt.Fprintf(
				os.Stderr,
				"[SDK PreFetch] Token acquired successfully\n",
			)
			state.clearFailure()
			return nil
		}
		lastErr = err
		fmt.Fprintf(
			os.Stderr,
			"[SDK PreFetch] Attempt %d failed: %v\n",
			attempt+1,
			err,
		)
	}

	state.recordFailure()
	return fmt.Errorf("PreFetch failed after 5 attempts: %w", lastErr)
}

// ClearFailure clears the failure throttle state, allowing retries.
func (state *OAuth2State) ClearFailure() {
	state.clearFailure()
}

func (state *OAuth2State) clearFailure() {
	state.failureMu.Lock()
	defer state.failureMu.Unlock()
	state.lastFailedAt = time.Time{}
}

func (state *OAuth2State) recordFailure() {
	state.failureMu.Lock()
	defer state.failureMu.Unlock()
	state.lastFailedAt = time.Now()
}

func (state *OAuth2State) isThrottled() bool {
	state.failureMu.Lock()
	defer state.failureMu.Unlock()
	return !state.lastFailedAt.IsZero() &&
		time.Since(state.lastFailedAt) < FailureCooldown
}

func (state *OAuth2State) GetToken(cfg *RequestConfig) (string, error) {
	// Log SDK version once on first token fetch
	sdkVersionLogged.Do(func() {
		fmt.Fprintf(
			os.Stderr,
			"[SDK] serval-go %s | TokenEndpoint: %s\n",
			SDK_VERSION,
			state.authPath,
		)
	})

	// Check throttle before attempting
	if state.isThrottled() {
		return "", fmt.Errorf(
			"token fetch throttled after recent failure (cooldown: %v)",
			FailureCooldown,
		)
	}

	tokenInfo := state.find(cfg)

	valid := tokenInfo.validToken()
	if valid != "" {
		return valid, nil
	}

	tokenInfo.reqTokenMutex.Lock()
	defer tokenInfo.reqTokenMutex.Unlock()

	// Double-check after acquiring lock (another goroutine might have refreshed)
	valid = tokenInfo.validToken()
	if valid != "" {
		return valid, nil
	}

	tokenInfo.reset()

	authUrl, err := state.authUrl(cfg)
	if err != nil {
		return "", err
	}

	// Create form body with grant_type (required by /v2/auth/token)
	formBody := strings.NewReader("grant_type=client_credentials")
	oAuthReq, err := http.NewRequestWithContext(
		cfg.Context,
		http.MethodPost,
		authUrl,
		formBody,
	)
	if err != nil {
		state.recordFailure()
		return "", fmt.Errorf(
			"requestconfig: failed to create OAuth2 token request: %w",
			err,
		)
	}

	encoded := base64.StdEncoding.EncodeToString(
		fmt.Appendf(nil, "%s:%s", cfg.ClientID, cfg.ClientSecret),
	)
	oAuthReq.Header.Set("Authorization", fmt.Sprintf("Basic %s", encoded))
	oAuthReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	handler := cfg.HTTPClient.Do
	if cfg.CustomHTTPDoer != nil {
		handler = cfg.CustomHTTPDoer.Do
	}

	oauthRes, err := handler(oAuthReq)
	if err != nil {
		state.recordFailure()
		return "", fmt.Errorf(
			"requestconfig: OAuth2 token request failed: %w",
			err,
		)
	}
	defer oauthRes.Body.Close()

	if oauthRes.StatusCode != http.StatusOK {
		state.recordFailure()
		body, _ := io.ReadAll(oauthRes.Body)
		return "", fmt.Errorf(
			"requestconfig: OAuth2 token request returned status %d: %s",
			oauthRes.StatusCode,
			string(body),
		)
	}

	contents, err := io.ReadAll(oauthRes.Body)
	if err != nil {
		state.recordFailure()
		return "", fmt.Errorf(
			"requestconfig: failed to read OAuth2 token response: %w",
			err,
		)
	}

	err = json.Unmarshal(contents, tokenInfo)
	if err != nil {
		state.recordFailure()
		return "", fmt.Errorf(
			"requestconfig: failed to parse OAuth2 token response: %w",
			err,
		)
	}

	if tokenInfo.AccessToken == "" {
		state.recordFailure()
		return "", fmt.Errorf(
			"requestconfig: OAuth2 token response missing access_token",
		)
	}

	tokenInfo.expiresAt = time.Now().
		Add(time.Duration(tokenInfo.ExpiresIn) * time.Second)
	state.clearFailure()

	return tokenInfo.AccessToken, nil
}

// fetchTokenDirect fetches a token directly without going through RequestConfig.
// Used by PreFetch to avoid circular dependencies.
func (state *OAuth2State) fetchTokenDirect(
	ctx context.Context,
	baseURL *url.URL,
	clientID, clientSecret string,
	httpClient *http.Client,
	tokenInfo *oAuthTokenInfo,
) error {
	tokenInfo.reqTokenMutex.Lock()
	defer tokenInfo.reqTokenMutex.Unlock()

	// Check again after acquiring lock
	if tokenInfo.validToken() != "" {
		return nil
	}

	tokenInfo.reset()

	authUrl, err := baseURL.Parse(strings.TrimLeft(state.authPath, "/"))
	if err != nil {
		return fmt.Errorf("failed to parse token URL: %w", err)
	}

	formBody := strings.NewReader("grant_type=client_credentials")
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		authUrl.String(),
		formBody,
	)
	if err != nil {
		return fmt.Errorf("failed to create token request: %w", err)
	}

	encoded := base64.StdEncoding.EncodeToString(
		[]byte(clientID + ":" + clientSecret),
	)
	req.Header.Set("Authorization", "Basic "+encoded)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("token request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read token response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"token request returned %d: %s",
			resp.StatusCode,
			string(body),
		)
	}

	err = json.Unmarshal(body, tokenInfo)
	if err != nil {
		return fmt.Errorf("failed to parse token response: %w", err)
	}

	if tokenInfo.AccessToken == "" {
		return fmt.Errorf("token response missing access_token")
	}

	tokenInfo.expiresAt = time.Now().
		Add(time.Duration(tokenInfo.ExpiresIn) * time.Second)
	return nil
}

func (state *OAuth2State) find(cfg *RequestConfig) *oAuthTokenInfo {
	return state.findByCredentials(cfg.ClientID, cfg.ClientSecret)
}

func (state *OAuth2State) findByCredentials(
	clientID, clientSecret string,
) *oAuthTokenInfo {
	key := oAuth2ClientCredentials{clientID, clientSecret}
	resp, ok := state.entries.Load(key)
	if !ok {
		resp, _ = state.entries.LoadOrStore(key, new(oAuthTokenInfo))
	}
	return resp.(*oAuthTokenInfo)
}

func (state *OAuth2State) authUrl(cfg *RequestConfig) (string, error) {
	authUrl, err := cfg.BaseURL.Parse(strings.TrimLeft(state.authPath, "/"))
	if err != nil {
		err = fmt.Errorf(
			"requestconfig: failed to parse OAuth2 token URL: %w",
			err,
		)
		return "", err
	}
	return authUrl.String(), nil
}
