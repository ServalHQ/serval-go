// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package serval_test

import (
	"context"
	"os"
	"testing"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/internal/testutil"
	"github.com/ServalHQ/serval-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := serval.NewClient(
		option.WithBaseURL(baseURL),
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
	)
	accessPolicy, err := client.AccessPolicies.New(context.TODO(), serval.AccessPolicyNewParams{
		Name: serval.String("Example Access Policy"),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", accessPolicy.ID)
}
