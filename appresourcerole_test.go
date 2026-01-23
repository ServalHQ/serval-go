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

func TestAppResourceRoleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.AppResourceRoles.New(context.TODO(), serval.AppResourceRoleNewParams{
		AccessPolicyID: serval.String("accessPolicyId"),
		Description:    serval.String("description"),
		ExternalData:   serval.String("externalData"),
		ExternalID:     serval.String("externalId"),
		Name:           serval.String("name"),
		ProvisioningMethod: serval.AppResourceRoleNewParamsProvisioningMethod{
			BuiltinWorkflow: map[string]any{},
			CustomWorkflow: serval.AppResourceRoleNewParamsProvisioningMethodCustomWorkflow{
				DeprovisionWorkflowID: serval.String("deprovisionWorkflowId"),
				ProvisionWorkflowID:   serval.String("provisionWorkflowId"),
			},
			LinkedRoles: serval.AppResourceRoleNewParamsProvisioningMethodLinkedRoles{
				LinkedRoleIDs: []string{"string"},
			},
			Manual: serval.AppResourceRoleNewParamsProvisioningMethodManual{
				Assignees: []serval.AppResourceRoleNewParamsProvisioningMethodManualAssignee{{
					AssigneeID:   serval.String("assigneeId"),
					AssigneeType: "MANUAL_PROVISIONING_ASSIGNEE_TYPE_UNSPECIFIED",
				}},
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

func TestAppResourceRoleGet(t *testing.T) {
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
	_, err := client.AppResourceRoles.Get(context.TODO(), "id")
	if err != nil {
		var apierr *serval.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAppResourceRoleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.AppResourceRoles.Update(
		context.TODO(),
		"id",
		serval.AppResourceRoleUpdateParams{
			AccessPolicyID: serval.String("accessPolicyId"),
			Description:    serval.String("description"),
			ExternalData:   serval.String("externalData"),
			ExternalID:     serval.String("externalId"),
			Name:           serval.String("name"),
			ProvisioningMethod: serval.AppResourceRoleUpdateParamsProvisioningMethod{
				BuiltinWorkflow: map[string]any{},
				CustomWorkflow: serval.AppResourceRoleUpdateParamsProvisioningMethodCustomWorkflow{
					DeprovisionWorkflowID: serval.String("deprovisionWorkflowId"),
					ProvisionWorkflowID:   serval.String("provisionWorkflowId"),
				},
				LinkedRoles: serval.AppResourceRoleUpdateParamsProvisioningMethodLinkedRoles{
					LinkedRoleIDs: []string{"string"},
				},
				Manual: serval.AppResourceRoleUpdateParamsProvisioningMethodManual{
					Assignees: []serval.AppResourceRoleUpdateParamsProvisioningMethodManualAssignee{{
						AssigneeID:   serval.String("assigneeId"),
						AssigneeType: "MANUAL_PROVISIONING_ASSIGNEE_TYPE_UNSPECIFIED",
					}},
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

func TestAppResourceRoleListWithOptionalParams(t *testing.T) {
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
	_, err := client.AppResourceRoles.List(context.TODO(), serval.AppResourceRoleListParams{
		AppInstanceID: serval.String("appInstanceId"),
		PageSize:      serval.Int(0),
		PageToken:     serval.String("pageToken"),
		ResourceID:    serval.String("resourceId"),
		TeamID:        serval.String("teamId"),
	})
	if err != nil {
		var apierr *serval.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAppResourceRoleDelete(t *testing.T) {
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
	_, err := client.AppResourceRoles.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *serval.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
