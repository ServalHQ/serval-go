// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package serval_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/internal/testutil"
	"github.com/ServalHQ/serval-go/option"
)

func TestAppResourceEntitlementNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.AppResourceEntitlements.New(context.TODO(), serval.AppResourceEntitlementNewParams{
		AccessPolicyID: serval.String("accessPolicyId"),
		Description:    serval.String("description"),
		ExternalData:   serval.String("externalData"),
		ExternalID:     serval.String("externalId"),
		Name:           serval.String("name"),
		ProvisioningMethod: serval.AppResourceEntitlementNewParamsProvisioningMethodUnion{
			OfBuiltinWorkflow: &serval.AppResourceEntitlementNewParamsProvisioningMethodBuiltinWorkflow{
				BuiltinWorkflow: map[string]any{},
			},
		},
		RequestsEnabled: serval.Bool(true),
		ResourceID:      serval.String("resourceId"),
	})
	if err != nil {
		var apierr *serval.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAppResourceEntitlementGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.AppResourceEntitlements.Get(context.TODO(), "id")
	if err != nil {
		var apierr *serval.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAppResourceEntitlementUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.AppResourceEntitlements.Update(
		context.TODO(),
		"id",
		serval.AppResourceEntitlementUpdateParams{
			AccessPolicyID: serval.String("accessPolicyId"),
			Description:    serval.String("description"),
			ExternalData:   serval.String("externalData"),
			ExternalID:     serval.String("externalId"),
			Name:           serval.String("name"),
			ProvisioningMethod: serval.AppResourceEntitlementUpdateParamsProvisioningMethodUnion{
				OfBuiltinWorkflow: &serval.AppResourceEntitlementUpdateParamsProvisioningMethodBuiltinWorkflow{
					BuiltinWorkflow: map[string]any{},
				},
			},
			RequestsEnabled: serval.Bool(true),
		},
	)
	if err != nil {
		var apierr *serval.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAppResourceEntitlementDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.AppResourceEntitlements.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *serval.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
