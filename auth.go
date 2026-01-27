// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.
// NOTE: PreFetch is a manual addition - preserve on regeneration.

package serval

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/ServalHQ/serval-go/internal/requestconfig"
)

// PreFetchToken fetches an OAuth token proactively before parallel operations begin.
// This should be called once at startup (e.g., in Terraform provider Configure)
// to avoid thundering herd when many parallel operations all need tokens.
// It retries with exponential backoff on failure.
func PreFetchToken(
	ctx context.Context,
	baseURL *url.URL,
	clientID, clientSecret string,
	httpClient *http.Client,
) error {
	oauth2State := requestconfig.OAuth2Cache["/v2/auth/token"]
	if oauth2State == nil {
		return nil
	}
	return oauth2State.PreFetch(
		ctx,
		baseURL,
		clientID,
		clientSecret,
		httpClient,
	)
}

// ClearTokenFailure clears the OAuth2 failure throttle state, allowing retries.
func ClearTokenFailure() {
	oauth2State := requestconfig.OAuth2Cache["/v2/auth/token"]
	if oauth2State != nil {
		oauth2State.ClearFailure()
	}
}

// GetToken returns a valid OAuth token, fetching one if necessary.
// This uses the SDK's shared token cache with throttling.
func GetToken(
	ctx context.Context,
	baseURL *url.URL,
	clientID, clientSecret string,
	httpClient *http.Client,
) (string, error) {
	// #region agent log
	fmt.Fprintf(
		os.Stderr,
		"[serval.GetToken] called, baseURL=%v, clientID=%q\n",
		baseURL,
		clientID,
	)
	// #endregion

	oauth2State := requestconfig.OAuth2Cache["/v2/auth/token"]
	if oauth2State == nil {
		// #region agent log
		fmt.Fprintf(os.Stderr, "[serval.GetToken] ERROR: oauth2State is nil\n")
		// #endregion
		return "", fmt.Errorf("OAuth2Cache not initialized")
	}

	// Build a minimal RequestConfig to get the token
	req, _ := http.NewRequestWithContext(ctx, "GET", baseURL.String(), nil)
	cfg := &requestconfig.RequestConfig{
		Context:      ctx,
		Request:      req,
		BaseURL:      baseURL,
		HTTPClient:   httpClient,
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	token, err := oauth2State.GetToken(cfg)
	// #region agent log
	fmt.Fprintf(
		os.Stderr,
		"[serval.GetToken] oauth2State.GetToken returned token=%d bytes, err=%v\n",
		len(token),
		err,
	)
	// #endregion
	return token, err
}
