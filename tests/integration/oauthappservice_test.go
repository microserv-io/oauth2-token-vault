//go:build integration

package integration

import (
	"context"
	"fmt"
	"github.com/microserv-io/oauth-credentials-server/pkg/proto/oauthcredentials/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
)

func TestGetOAuthByProvider(t *testing.T) {
	// Set up the gRPC connection
	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%s", ServerPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := oauthcredentials.NewOAuthServiceClient(conn)

	providerClient := oauthcredentials.NewOAuthProviderServiceClient(conn)

	providerResp, err := providerClient.CreateProvider(context.Background(), &oauthcredentials.CreateProviderRequest{
		Name: "Test Provider",
	})

	if err != nil {
		t.Fatalf("CreateProvider failed: %v", err)
	}

	// Create a new provider
	req := &oauthcredentials.GetOAuthByProviderRequest{
		Provider: providerResp.GetOauthProvider().GetName(),
	}

	resp, err := client.GetOAuthByProvider(context.Background(), req)

	log.Print(resp)

	if err != nil {
		t.Fatalf("CreateProvider failed: %v", err)
	}

	assert.NotNil(t, resp)
	assert.Equal(t, "Test Provider", resp.GetOauth().GetProvider())
}
