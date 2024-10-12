package oauthapp

import (
	"context"
	"fmt"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/oauthapp"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/provider"
	oauth3 "github.com/microserv-io/oauth-credentials-server/internal/domain/oauth2"
	"golang.org/x/oauth2"
	"log/slog"
	"net/url"
	"strconv"
)

type OAuthAppRepository interface {
	oauthapp.Repository
}

type ProviderRepository interface {
	provider.Repository
}

type TokenSourceFactory interface {
	oauth3.TokenSourceFactory
}

type Service struct {
	oauthAppRepository OAuthAppRepository
	providerRepository ProviderRepository
	tokenSourceFactory TokenSourceFactory

	logger *slog.Logger
}

func NewService(
	oauthAppRepository OAuthAppRepository,
	providerRepository ProviderRepository,
	tokenSourceFactory TokenSourceFactory,
	logger *slog.Logger,
) *Service {
	return &Service{
		oauthAppRepository: oauthAppRepository,
		providerRepository: providerRepository,
		tokenSourceFactory: tokenSourceFactory,
		logger:             logger,
	}
}

// ListOAuthsForOwner lists all oauth apps for a given owner
func (s *Service) ListOAuthAppsForOwner(ctx context.Context, ownerID string) (*ListOAuthAppsForOwnerResponse, error) {
	oauthApps, err := s.oauthAppRepository.ListForOwner(ctx, ownerID)
	if err != nil {
		return nil, fmt.Errorf("failed to list oauth apps for owner: %w", err)
	}

	response := &ListOAuthAppsForOwnerResponse{
		Apps: make([]*OAuthApp, 0, len(oauthApps)),
	}

	for _, app := range oauthApps {
		response.Apps = append(response.Apps, &OAuthApp{
			ID:         strconv.Itoa(int(app.ID)),
			OwnerID:    app.OwnerID,
			ProviderID: app.Provider,
			Scopes:     app.Scopes,
		})
	}

	return response, nil
}

func (s *Service) GetOAuthForProviderAndOwner(ctx context.Context, providerID string, ownerID string) (*GetOAuthForProviderAndOwnerResponse, error) {
	oauthApp, err := s.oauthAppRepository.Find(ctx, ownerID, providerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get oauth app by id: %w", err)
	}

	return &GetOAuthForProviderAndOwnerResponse{
		App: &OAuthApp{
			ID:         strconv.Itoa(int(oauthApp.ID)),
			OwnerID:    oauthApp.OwnerID,
			ProviderID: oauthApp.Provider,
			Scopes:     oauthApp.Scopes,
		},
	}, nil
}

func (s *Service) CreateAuthorizationURLForProvider(ctx context.Context, providerID string, scopes []string, state string) (*CreateAuthorizationURLForProviderResponse, error) {
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

	return &CreateAuthorizationURLForProviderResponse{
		URL: parsedURL,
	}, nil
}

func (s *Service) RetrieveAccessToken(ctx context.Context, providerID string, ownerID string) (*RetrieveAccessTokenResponse, error) {
	oauthApp, err := s.oauthAppRepository.Find(ctx, ownerID, providerID)

	if err != nil {
		return nil, fmt.Errorf("failed to find oauth app by id: %w", err)
	}

	providerObj, err := s.providerRepository.FindByName(ctx, providerID)
	if err != nil {
		return nil, fmt.Errorf("failed to find provider by name: %w", err)
	}

	tokenSource := s.tokenSourceFactory.NewTokenSource(ctx, providerObj, oauthApp)

	newToken, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve access token: %w", err)
	}

	oauthApp.AccessToken = newToken.AccessToken
	oauthApp.RefreshToken = newToken.RefreshToken
	oauthApp.ExpiresAt = newToken.Expiry

	err = s.oauthAppRepository.UpdateByID(ctx, oauthApp.ID, func(app *oauthapp.OAuthApp) error {
		app.AccessToken = newToken.AccessToken
		app.RefreshToken = newToken.RefreshToken
		app.ExpiresAt = newToken.Expiry
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update oauth app token: %w", err)
	}

	return &RetrieveAccessTokenResponse{
		AccessToken: newToken.AccessToken,
	}, nil
}
