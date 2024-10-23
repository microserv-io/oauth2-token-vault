package oauthapp

import (
	"context"
	"fmt"
	"github.com/microserv-io/oauth2-token-vault/internal/domain"
	"github.com/microserv-io/oauth2-token-vault/internal/domain/models/oauthapp"
	"github.com/microserv-io/oauth2-token-vault/internal/domain/models/provider"
	"github.com/microserv-io/oauth2-token-vault/internal/domain/oauth2"
	"log/slog"
	"strconv"
)

type OAuthAppRepository interface {
	oauthapp.Repository
}

type ProviderRepository interface {
	provider.Repository
}

type TokenSourceFactory interface {
	oauth2.TokenSourceFactory
}

type Encryptor interface {
	domain.Encryptor
}

type Service struct {
	oauthAppRepository OAuthAppRepository
	providerRepository ProviderRepository
	tokenSourceFactory TokenSourceFactory
	encryptor          Encryptor

	logger *slog.Logger
}

func NewService(
	oauthAppRepository OAuthAppRepository,
	providerRepository ProviderRepository,
	tokenSourceFactory TokenSourceFactory,
	encryptor Encryptor,
	logger *slog.Logger,
) *Service {
	return &Service{
		oauthAppRepository: oauthAppRepository,
		providerRepository: providerRepository,
		tokenSourceFactory: tokenSourceFactory,
		encryptor:          encryptor,
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

func (s *Service) RetrieveAccessToken(ctx context.Context, providerID string, ownerID string) (*RetrieveAccessTokenResponse, error) {
	oauthApp, err := s.oauthAppRepository.Find(ctx, ownerID, providerID)

	if err != nil {
		return nil, fmt.Errorf("failed to find oauth app by id: %w", err)
	}

	providerObj, err := s.providerRepository.FindByName(ctx, providerID)
	if err != nil {
		return nil, fmt.Errorf("failed to find provider by name: %w", err)
	}

	clientSecret, err := s.encryptor.Decrypt(providerObj.ClientSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt client secret: %w", err)
	}

	accessToken, err := s.encryptor.Decrypt(oauthApp.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt access token: %w", err)
	}

	refreshToken, err := s.encryptor.Decrypt(oauthApp.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt refresh token: %w", err)
	}

	tokenSource, err := s.tokenSourceFactory.NewTokenSource(ctx, &oauth2.TokenSourceConfig{
		ClientID:     providerObj.ClientID,
		ClientSecret: clientSecret,
		AuthURL:      providerObj.AuthURL,
		TokenURL:     providerObj.TokenURL,
		RedirectURL:  providerObj.RedirectURL,
		Scopes:       oauthApp.Scopes,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    oauthApp.TokenType,
		ExpiresAt:    oauthApp.ExpiresAt,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create token source: %w", err)
	}

	newToken, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve access token: %w", err)
	}

	if err = s.oauthAppRepository.UpdateByID(ctx, oauthApp.ID, func(app *oauthapp.OAuthApp) error {
		accessToken, err := s.encryptor.Encrypt(newToken.AccessToken)
		if err != nil {
			return fmt.Errorf("failed to encrypt access token: %w", err)
		}

		refreshToken, err := s.encryptor.Encrypt(newToken.RefreshToken)
		if err != nil {
			return fmt.Errorf("failed to encrypt refresh token: %w", err)
		}

		app.AccessToken = accessToken
		app.RefreshToken = refreshToken
		app.ExpiresAt = newToken.Expiry
		return nil
	}); err != nil {
		return nil, fmt.Errorf("failed to update oauth app: %w", err)
	}

	return &RetrieveAccessTokenResponse{
		AccessToken: newToken.AccessToken,
	}, nil
}
