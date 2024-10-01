package usecase

import (
	"context"
	"fmt"
	"github.com/microserv-io/oauth-credentials-server/internal/domain"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/oauthapp"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/provider"
	"time"
)

type Credential struct {
	TokenType    string
	ExpiresAt    time.Time
	AccessToken  string
	RefreshToken string
}

type GetCredentialsUseCase struct {
	repository         oauthapp.Repository
	providerRepository provider.Repository
	tokenSourceFactory domain.TokenSourceFactory
}

func NewGetCredentialsUseCase(repository oauthapp.Repository, providerRepository provider.Repository, tokenSourceFactory domain.TokenSourceFactory) *GetCredentialsUseCase {
	return &GetCredentialsUseCase{
		repository:         repository,
		providerRepository: providerRepository,
		tokenSourceFactory: tokenSourceFactory,
	}
}

func (u *GetCredentialsUseCase) Execute(ctx context.Context, id string, ownerID string) (Credential, error) {
	oauthApp, err := u.repository.Find(ctx, ownerID, id)
	if err != nil {
		return Credential{}, fmt.Errorf("failed to find oauth app: %w", err)
	}

	p, err := u.providerRepository.FindByName(ctx, oauthApp.Provider)
	if err != nil {
		return Credential{}, fmt.Errorf("failed to find provider: %w", err)
	}

	tokenSource := u.tokenSourceFactory.NewTokenSource(ctx, *p, *oauthApp)
	newToken, err := tokenSource.Token()

	if err != nil {
		return Credential{}, fmt.Errorf("failed to refresh token: %w", err)
	}

	if err := u.repository.Update(ctx, oauthApp.ID, func(app *oauthapp.OAuthApp) error {
		oauthApp.AccessToken = newToken.AccessToken
		oauthApp.RefreshToken = newToken.RefreshToken
		oauthApp.ExpiresAt = newToken.Expiry
		return nil
	}); err != nil {
		return Credential{}, fmt.Errorf("failed to update token: %w", err)
	}

	return Credential{
		TokenType:    oauthApp.TokenType,
		ExpiresAt:    oauthApp.ExpiresAt,
		AccessToken:  oauthApp.AccessToken,
		RefreshToken: oauthApp.RefreshToken,
	}, nil
}
