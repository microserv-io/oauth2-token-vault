package usecase

import (
	"context"
	"github.com/microserv-io/oauth/internal/domain/oauthapp"
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
	providerRepository ProviderRepository
}

func NewGetCredentialsUseCase(repository oauthapp.Repository, providerRepository ProviderRepository) *GetCredentialsUseCase {
	return &GetCredentialsUseCase{
		repository:         repository,
		providerRepository: providerRepository,
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

	config := oauth2.Config{
		ClientID:     providerConfig.ClientID,
		ClientSecret: providerConfig.ClientSecret,
		Scopes:       providerConfig.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  providerConfig.AuthURL,
			TokenURL: providerConfig.TokenURL,
		},
		RedirectURL: providerConfig.RedirectURL,
	}

	token := oauth2.Token{
		AccessToken:  oauthApp.AccessToken,
		TokenType:    "",
		RefreshToken: oauthApp.RefreshToken,
		Expiry:       oauthApp.ExpiresAt,
	}

	_, err = config.TokenSource(oauth2.NoContext, &token).Token()
	if err != nil {
		return Credential{}, err
	}

	oauthApp.AccessToken = token.AccessToken
	oauthApp.RefreshToken = token.RefreshToken
	oauthApp.ExpiresAt = token.Expiry

	if err := u.repository.Update(ctx, oauthApp); err != nil {
		return Credential{}, err
	}

	return Credential{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}
