package providerservice

import (
	"context"
	"fmt"
	"github.com/microserv-io/oauth-credentials-server/internal/app/oauthapp"
	"github.com/microserv-io/oauth-credentials-server/internal/app/provider"
	"github.com/microserv-io/oauth-credentials-server/pkg/proto/oauthcredentials/v1"
)

var _ oauthcredentials.OAuthProviderServiceServer = &Service{}

type Service struct {
	oauthcredentials.UnimplementedOAuthProviderServiceServer
	providerService *provider.Service
	oauthappService *oauthapp.Service
}

func NewService(
	providerService *provider.Service,
	oauthAppService *oauthapp.Service,
) *Service {
	return &Service{
		providerService: providerService,
		oauthappService: oauthAppService,
	}
}

func (s Service) ListProviders(_ *oauthcredentials.ListProvidersRequest, stream oauthcredentials.OAuthProviderService_ListProvidersServer) error {
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

func (s Service) CreateProvider(ctx context.Context, oauthProvider *oauthcredentials.CreateProviderRequest) (*oauthcredentials.CreateProviderResponse, error) {
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
func (s Service) UpdateProvider(ctx context.Context, oauthProvider *oauthcredentials.UpdateProviderRequest) (*oauthcredentials.UpdateProviderResponse, error) {
	resp, err := s.providerService.Update(ctx, oauthProvider.Name, &provider.UpdateInput{
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

func (s Service) DeleteProvider(ctx context.Context, oauthProvider *oauthcredentials.DeleteProviderRequest) (*oauthcredentials.DeleteProviderResponse, error) {
	err := s.providerService.DeleteProvider(ctx, oauthProvider.GetId())
	if err != nil {
		return nil, fmt.Errorf("failed to delete provider: %w", err)
	}

	return &oauthcredentials.DeleteProviderResponse{}, nil
}
