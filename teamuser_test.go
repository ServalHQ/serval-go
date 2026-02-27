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

func TestTeamUserNew(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := serval.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Teams.Users.New(
		context.TODO(),
		"team_id",
		serval.TeamUserNewParams{
			Role:   serval.TeamUserNewParamsRoleTeamUserRoleUnspecified,
			UserID: "userId",
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

func TestTeamUserGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := serval.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Teams.Users.Get(
		context.TODO(),
		"user_id",
		serval.TeamUserGetParams{
			TeamID: "team_id",
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

func TestTeamUserUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := serval.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Teams.Users.Update(
		context.TODO(),
		"user_id",
		serval.TeamUserUpdateParams{
			TeamID: "team_id",
			Role:   serval.TeamUserUpdateParamsRoleTeamUserRoleUnspecified,
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

func TestTeamUserListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := serval.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Teams.Users.List(
		context.TODO(),
		"team_id",
		serval.TeamUserListParams{
			PageSize:  serval.Int(0),
			PageToken: serval.String("pageToken"),
			UserID:    serval.String("userId"),
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

func TestTeamUserDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := serval.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Teams.Users.Delete(
		context.TODO(),
		"user_id",
		serval.TeamUserDeleteParams{
			TeamID: "team_id",
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
