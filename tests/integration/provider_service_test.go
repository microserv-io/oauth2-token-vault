//go:build integration

package integration

import (
	"context"
	"fmt"
	"github.com/microserv-io/oauth-credentials-server/pkg/proto/oauthcredentials/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

// TODO(roccolangeweg): Currently the ordering of the tests matter. Create must be executed first, delete must be executed last. Tests need to be standalone (e.g. unique IDs)

func TestCreateProvider(t *testing.T) {
	// Set up the gRPC connection
	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%s", ServerPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := oauthcredentials.NewOAuthProviderServiceClient(conn)

	// Create a new provider
	req := &oauthcredentials.CreateProviderRequest{
		Name: "Test Provider",
		// Add other necessary fields here
	}

	resp, err := client.CreateProvider(context.Background(), req)
	if err != nil {
		t.Fatalf("CreateProvider failed: %v", err)
	}

	assert.NotNil(t, resp)
	assert.Equal(t, "Test Provider", resp.GetOauthProvider().GetName())
}

func TestListProviders(t *testing.T) {
	// Set up the gRPC connection
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", ServerPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := oauthcredentials.NewOAuthProviderServiceClient(conn)

	// List all providers
	req := &oauthcredentials.ListProvidersRequest{}

	stream, err := client.ListProviders(context.Background(), req)
	if err != nil {
		t.Fatalf("ListProviders failed: %v", err)
	}

	var providers []*oauthcredentials.OAuthProvider
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		for _, p := range resp.GetOauthProviders() {
			providers = append(providers, p)
		}
	}

	assert.NotEmpty(t, providers)
}

func TestUpdateProvider(t *testing.T) {
	// Set up the gRPC connection
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", ServerPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := oauthcredentials.NewOAuthProviderServiceClient(conn)

	// Update a provider
	req := &oauthcredentials.UpdateProviderRequest{
		Name:    "Test Provider",
		AuthUrl: "https://example.com/auth",
		// Add other necessary fields here
	}

	resp, err := client.UpdateProvider(context.Background(), req)
	if err != nil {
		t.Fatalf("UpdateProvider failed: %v", err)
	}

	assert.NotNil(t, resp)
	assert.Equal(t, "Test Provider", resp.GetOauthProvider().GetName())
	assert.Equal(t, "https://example.com/auth", resp.GetOauthProvider().GetAuthUrl())
}

func TestDeleteProvider(t *testing.T) {
	// Set up the gRPC connection
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", ServerPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := oauthcredentials.NewOAuthProviderServiceClient(conn)

	// Delete a provider
	req := &oauthcredentials.DeleteProviderRequest{
		Id: "Test Provider",
	}

	resp, err := client.DeleteProvider(context.Background(), req)
	if err != nil {
		t.Fatalf("DeleteProvider failed: %v", err)
	}

	assert.NotNil(t, resp)

	// Check if the provider was deleted
	listProviderRequest := &oauthcredentials.ListProvidersRequest{}

	stream, err := client.ListProviders(context.Background(), listProviderRequest)
	if err != nil {
		t.Fatalf("ListProviders failed: %v", err)
	}

	var providers []*oauthcredentials.OAuthProvider
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		for _, p := range resp.GetOauthProviders() {
			providers = append(providers, p)
		}
	}

	assert.Empty(t, providers)
}
