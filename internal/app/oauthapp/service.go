package oauthapp

import (
	"context"
	"fmt"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/oauthapp"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/provider"
	"golang.org/x/oauth2"
	"log/slog"
	"net/url"
)

type OAuthAppRepository interface {
	ListForOwner(ctx context.Context, ownerID string) ([]*oauthapp.OAuthApp, error)
	Find(ctx context.Context, ownerID string, providerID string) (*oauthapp.OAuthApp, error)
}

type ProviderRepository interface {
	FindByName(ctx context.Context, name string) (*provider.Provider, error)
}

type Service struct {
	oauthAppRepository OAuthAppRepository
	providerRepository ProviderRepository

	logger *slog.Logger
}

func NewOAuthAppService(oauthAppRepository OAuthAppRepository, providerRepository ProviderRepository, logger *slog.Logger) *Service {
	return &Service{
		oauthAppRepository: oauthAppRepository,
		providerRepository: providerRepository,
		logger:             logger,
	}
}

// ListOAuthsForOwner lists all oauth apps for a given owner
func (s *Service) ListOAuthsForOwner(ctx context.Context, ownerID string) ([]*OAuthApp, error) {
	oauthApps, err := s.oauthAppRepository.ListForOwner(ctx, ownerID)
	if err != nil {
		return nil, fmt.Errorf("failed to list oauth apps for owner: %w", err)
	}

	apps := make([]*OAuthApp, 0, len(oauthApps))

	for _, app := range oauthApps {
		apps = append(apps, &OAuthApp{
			ID:         app.ID,
			OwnerID:    app.OwnerID,
			ProviderID: app.Provider,
			Scopes:     app.Scopes,
		})
	}

	return apps, nil
}

func (s *Service) GetOAuthForProviderAndOwner(ctx context.Context, providerID string, ownerID string) (*OAuthApp, error) {
	oauthApp, err := s.oauthAppRepository.Find(ctx, ownerID, providerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get oauth app by id: %w", err)
	}

	return &OAuthApp{
		ID:         oauthApp.ID,
		OwnerID:    oauthApp.OwnerID,
		ProviderID: oauthApp.Provider,
		Scopes:     oauthApp.Scopes,
	}, nil
}

func (s *Service) CreateAuthorizationURLForProvider(ctx context.Context, providerID string, scopes []string, state string) (*AuthorizationURLResponse, error) {
	providerObj, err := s.providerRepository.FindByName(ctx, providerID)
	if err != nil {
		return nil, fmt.Errorf("failed to find provider by name: %w", err)
	}

	if len(scopes) == 0 {
		scopes = providerObj.Scopes
	}

	oauth2Config := oauth2.Config{
		ClientID:     providerObj.ClientID,
		ClientSecret: providerObj.ClientSecret,
		RedirectURL:  providerObj.RedirectURL,
		Scopes:       scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  providerObj.AuthURL,
			TokenURL: providerObj.TokenURL,
		},
	}

	redirectURL := oauth2Config.AuthCodeURL(state)
	if redirectURL == "" {
		return nil, fmt.Errorf("failed to create redirect url for provider: %s", providerID)
	}

	parsedURL, err := url.Parse(redirectURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse redirect url: %w", err)
	}

	return &AuthorizationURLResponse{
		URL: parsedURL,
	}, nil
}
