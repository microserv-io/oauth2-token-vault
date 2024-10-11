package v1

import (
	"context"
	"fmt"
	"github.com/microserv-io/oauth-credentials-server/internal/app/provider"
	"github.com/microserv-io/oauth-credentials-server/pkg/proto/oauthcredentials/v1"
)

var _ oauthcredentials.OAuthProviderServiceServer = &ProviderServiceGRPC{}

type ProviderService interface {
	ListProviders(ctx context.Context) (*provider.ListProvidersResponse, error)
	CreateProvider(ctx context.Context, input *provider.CreateInput, ownerID string) (*provider.CreateProviderResponse, error)
	UpdateProvider(ctx context.Context, name string, input *provider.UpdateInput) (*provider.UpdateProviderResponse, error)
	DeleteProvider(ctx context.Context, id string) error
	ExchangeAuthorizationCode(ctx context.Context, input *provider.ExchangeAuthorizationCodeInput) error
}

type ListProviderStream interface {
	oauthcredentials.OAuthProviderService_ListProvidersServer
}

type ProviderServiceGRPC struct {
	oauthcredentials.UnimplementedOAuthProviderServiceServer
	providerService ProviderService
}

func NewProviderServiceGRPC(
	providerService ProviderService,
) *ProviderServiceGRPC {
	return &ProviderServiceGRPC{
		providerService: providerService,
	}
}

func (s ProviderServiceGRPC) ListProviders(_ *oauthcredentials.ListProvidersRequest, stream oauthcredentials.OAuthProviderService_ListProvidersServer) error {
	resp, err := s.providerService.ListProviders(stream.Context())
	if err != nil {
		return fmt.Errorf("failed to list providers: %w", err)
	}

	var oauthProviders []*oauthcredentials.OAuthProvider

	for _, p := range resp.Providers {
		oauthProviders = append(oauthProviders, &oauthcredentials.OAuthProvider{
			Name:     p.Name,
			AuthUrl:  p.AuthURL,
			TokenUrl: p.TokenURL,
			Scopes:   p.Scopes,
			ClientId: p.ClientID,
		})
	}

	if err := stream.Send(&oauthcredentials.ListProvidersResponse{OauthProviders: oauthProviders}); err != nil {
		return fmt.Errorf("failed to send provider: %w", err)
	}

	return nil
}

func (s ProviderServiceGRPC) CreateProvider(ctx context.Context, oauthProvider *oauthcredentials.CreateProviderRequest) (*oauthcredentials.CreateProviderResponse, error) {
	resp, err := s.providerService.CreateProvider(ctx, &provider.CreateInput{
		Name:    oauthProvider.Name,
		AuthURL: oauthProvider.AuthUrl,
	}, "api")
	if err != nil {
		return nil, fmt.Errorf("failed to create provider: %w", err)
	}

	return &oauthcredentials.CreateProviderResponse{
		OauthProvider: &oauthcredentials.OAuthProvider{
			Name:     resp.Provider.Name,
			AuthUrl:  resp.Provider.AuthURL,
			TokenUrl: resp.Provider.TokenURL,
			Scopes:   resp.Provider.Scopes,
			ClientId: resp.Provider.ClientID,
		},
	}, nil
}
func (s ProviderServiceGRPC) UpdateProvider(ctx context.Context, oauthProvider *oauthcredentials.UpdateProviderRequest) (*oauthcredentials.UpdateProviderResponse, error) {
	resp, err := s.providerService.UpdateProvider(ctx, oauthProvider.Name, &provider.UpdateInput{
		ClientID:     oauthProvider.ClientId,
		ClientSecret: oauthProvider.ClientSecret,
		RedirectURI:  oauthProvider.RedirectUri,
		Scopes:       oauthProvider.Scopes,
		AuthURL:      oauthProvider.AuthUrl,
		TokenURL:     oauthProvider.TokenUrl,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update provider: %w", err)
	}

	return &oauthcredentials.UpdateProviderResponse{
		OauthProvider: &oauthcredentials.OAuthProvider{
			Name:     resp.Provider.Name,
			AuthUrl:  resp.Provider.AuthURL,
			TokenUrl: resp.Provider.TokenURL,
			Scopes:   resp.Provider.Scopes,
			ClientId: resp.Provider.ClientID,
		},
	}, nil
}

func (s ProviderServiceGRPC) DeleteProvider(ctx context.Context, oauthProvider *oauthcredentials.DeleteProviderRequest) (*oauthcredentials.DeleteProviderResponse, error) {
	err := s.providerService.DeleteProvider(ctx, oauthProvider.GetId())
	if err != nil {
		return nil, fmt.Errorf("failed to delete provider: %w", err)
	}

	return &oauthcredentials.DeleteProviderResponse{}, nil
}

func (s ProviderServiceGRPC) ExchangeAuthorizationCode(ctx context.Context, input *oauthcredentials.ExchangeAuthorizationCodeRequest) (*oauthcredentials.ExchangeAuthorizationCodeResponse, error) {
	if err := s.providerService.ExchangeAuthorizationCode(ctx, &provider.ExchangeAuthorizationCodeInput{
		Provider: input.GetProvider(),
		Code:     input.GetCode(),
	}); err != nil {
		return nil, fmt.Errorf("failed to exchange authorization code: %w", err)
	}

	return &oauthcredentials.ExchangeAuthorizationCodeResponse{}, nil
}
