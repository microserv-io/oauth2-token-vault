package v1

import (
	"context"
	"fmt"
	"github.com/microserv-io/oauth2-token-vault/internal/app/provider"
	"github.com/microserv-io/oauth2-token-vault/pkg/proto/oauthcredentials/v1"
)

var _ oauthcredentials.OAuthProviderServiceServer = &ProviderServiceGRPC{}

type ProviderService interface {
	ListProviders(ctx context.Context) (*provider.ListProvidersResponse, error)
	CreateProvider(ctx context.Context, input *provider.CreateProviderRequest, ownerID string) (*provider.CreateProviderResponse, error)
	UpdateProvider(ctx context.Context, name string, input *provider.UpdateProviderRequest) (*provider.UpdateProviderResponse, error)
	DeleteProvider(ctx context.Context, request *provider.DeleteProviderRequest) error
	ExchangeAuthorizationCode(ctx context.Context, input *provider.ExchangeAuthorizationCodeRequest) error
	GetAuthorizationURL(ctx context.Context, input *provider.GetAuthorizationURLRequest) (*provider.GetAuthorizationURLResponse, error)
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

func (s ProviderServiceGRPC) ListProviders(ctx context.Context, _ *oauthcredentials.ListProvidersRequest) (*oauthcredentials.ListProvidersResponse, error) {
	resp, err := s.providerService.ListProviders(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list providers: %w", err)
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

	return &oauthcredentials.ListProvidersResponse{
		OauthProviders: oauthProviders,
	}, nil
}

func (s ProviderServiceGRPC) CreateProvider(ctx context.Context, oauthProvider *oauthcredentials.CreateProviderRequest) (*oauthcredentials.CreateProviderResponse, error) {
	resp, err := s.providerService.CreateProvider(ctx, &provider.CreateProviderRequest{
		Name:         oauthProvider.Name,
		AuthURL:      oauthProvider.AuthUrl,
		TokenURL:     oauthProvider.TokenUrl,
		Scopes:       oauthProvider.Scopes,
		ClientID:     oauthProvider.ClientId,
		ClientSecret: oauthProvider.ClientSecret,
		RedirectURI:  oauthProvider.RedirectUri,
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
	resp, err := s.providerService.UpdateProvider(ctx, oauthProvider.Name, &provider.UpdateProviderRequest{
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
	err := s.providerService.DeleteProvider(ctx, &provider.DeleteProviderRequest{Name: oauthProvider.GetName()})
	if err != nil {
		return nil, fmt.Errorf("failed to delete provider: %w", err)
	}

	return &oauthcredentials.DeleteProviderResponse{}, nil
}

func (s ProviderServiceGRPC) GetAuthorizationURL(ctx context.Context, input *oauthcredentials.GetAuthorizationURLRequest) (*oauthcredentials.GetAuthorizationURLResponse, error) {
	resp, err := s.providerService.GetAuthorizationURL(ctx, &provider.GetAuthorizationURLRequest{
		Provider: input.GetProvider(),
		State:    input.GetState(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get authorization url: %w", err)
	}

	return &oauthcredentials.GetAuthorizationURLResponse{
		Url: resp.URL.String(),
	}, nil
}

func (s ProviderServiceGRPC) ExchangeAuthorizationCode(ctx context.Context, input *oauthcredentials.ExchangeAuthorizationCodeRequest) (*oauthcredentials.ExchangeAuthorizationCodeResponse, error) {
	if err := s.providerService.ExchangeAuthorizationCode(ctx, &provider.ExchangeAuthorizationCodeRequest{
		Provider: input.GetProvider(),
		Code:     input.GetCode(),
	}); err != nil {
		return nil, fmt.Errorf("failed to exchange authorization code: %w", err)
	}

	return &oauthcredentials.ExchangeAuthorizationCodeResponse{}, nil
}
