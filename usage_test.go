// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package serval_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/serval-go"
	"github.com/stainless-sdks/serval-go/internal/testutil"
	"github.com/stainless-sdks/serval-go/option"
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
		option.WithAPIKey("My API Key"),
	)
	accessPolicy, err := client.AccessPolicies.Get(context.TODO(), "REPLACE_ME")
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", accessPolicy.ID)
}
