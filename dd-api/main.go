package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	ctx := context.WithValue(
		context.Background(),
		datadog.ContextAPIKeys,
		map[string]datadog.APIKey{
			"apiKeyAuth": {
				Key: os.Getenv("DD_CLIENT_API_KEY"),
			},
			"appKeyAuth": {
				Key: os.Getenv("DD_CLIENT_APP_KEY"),
			},
		},
	)

	if site, ok := os.LookupEnv("DD_SITE"); ok {
		ctx = context.WithValue(
			ctx,
			datadog.ContextServerVariables,
			map[string]string{"site": site},
		)
	}

	configuration := datadog.NewConfiguration()

	api_client := datadog.NewAPIClient(configuration)
	resp, r, err := api_client.UsersApi.ListUsers(ctx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ListUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	email := *(*resp.Data)[0].Attributes.Email
	fmt.Printf("user id: %s\n", email)

	// response from `ListUsers`: UsersResponse
	response_content, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from UsersApi.ListUsers:\n%s\n", response_content)
}
