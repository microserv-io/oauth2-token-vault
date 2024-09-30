package usecase

import (
	"context"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/oauthapp"
	"golang.org/x/oauth2"
)

type ProviderConfig struct {
	ClientID     string
	ClientSecret string
	Scopes       []string
	AuthURL      string
	TokenURL     string
	RedirectURL  string
}

type ProviderRepository interface {
	Find(provider string) (*ProviderConfig, error)
}

type Credential struct {
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
	oauthApp, err := u.repository.Find(ctx, id, ownerID)
	if err != nil {
		return Credential{}, err
	}

	providerConfig, err := u.providerRepository.Find(oauthApp.Provider)
	if err != nil {
		return Credential{}, err
	}

	tokenSource := u.tokenSourceFactory.NewTokenSource(ctx, *p, *oauthApp)
	newToken, err := tokenSource.Token()

	if err != nil {
		return Credential{}, err
	}

	if err := u.repository.Update(ctx, oauthApp.ID, func(app *oauthapp.OAuthApp) error {
		oauthApp.AccessToken = newToken.AccessToken
		oauthApp.RefreshToken = newToken.RefreshToken
		oauthApp.ExpiresAt = newToken.Expiry
		return nil
	}); err != nil {
		return Credential{}, err
	}

	return Credential{
		TokenType:    oauthApp.TokenType,
		ExpiresAt:    oauthApp.ExpiresAt,
		AccessToken:  oauthApp.AccessToken,
		RefreshToken: oauthApp.RefreshToken,
	}, nil
}
