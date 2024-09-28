package usecase

import (
	"context"
	"fmt"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/oauthapp"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/provider"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/oauth2"
)

type ExchangeAuthorizationCodeUseCase struct {
	oauthAppRepository oauthapp.Repository
	providerRepository provider.Repository
	oauth2Client       oauth2.Client
}

func NewExchangeAuthorizationCodeUseCase(
	oauthAppRepository oauthapp.Repository,
	providerRepository provider.Repository,
) *ExchangeAuthorizationCodeUseCase {
	return &ExchangeAuthorizationCodeUseCase{
		oauthAppRepository: oauthAppRepository,
		providerRepository: providerRepository,
	}
}

func (u *ExchangeAuthorizationCodeUseCase) Execute(ctx context.Context, providerName string, ownerID string, code string) error {
	p, err := u.providerRepository.FindByName(ctx, providerName)
	if err != nil {
		return fmt.Errorf("could not find provider %s: %w", providerName, err)
	}

	oauth2Config := &oauth2.Config{
		ClientID:     p.ClientID,
		ClientSecret: p.ClientSecret,
		AuthURL:      p.AuthURL,
		TokenURL:     p.TokenURL,
		RedirectURL:  p.RedirectURL,
		Scopes:       p.Scopes,
	}
	token, err := u.oauth2Client.Exchange(ctx, oauth2Config, code)
	if err != nil {
		return fmt.Errorf("could not exchange code for token: %w", err)
	}

	oauthApp, err := u.oauthAppRepository.Find(ctx, providerName, ownerID)
	if err != nil {
		if err.Error() == "not found" {
			return u.createOAuthApp(ctx, providerName, ownerID, token, p.Scopes)
		}
		return fmt.Errorf("error trying to see if oauth app exists: %w", err)
	}

	return u.saveToken(ctx, oauthApp, token)
}

func (u *ExchangeAuthorizationCodeUseCase) createOAuthApp(ctx context.Context, providerName, ownerID string, token *oauth2.Token, scopes []string) error {
	oauthApp := &oauthapp.OAuthApp{
		Provider:     providerName,
		OwnerID:      ownerID,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresAt:    token.ExpiresAt,
		Scopes:       scopes,
	}

	if err := u.oauthAppRepository.Create(ctx, oauthApp); err != nil {
		return fmt.Errorf("could not create oauth app: %w", err)
	}
	return nil
}

func (u *ExchangeAuthorizationCodeUseCase) saveToken(ctx context.Context, oauthApp *oauthapp.OAuthApp, token *oauth2.Token) error {
	return u.oauthAppRepository.Update(ctx, oauthApp.ID, func(app *oauthapp.OAuthApp) error {
		app.AccessToken = token.AccessToken
		app.RefreshToken = token.RefreshToken
		app.ExpiresAt = token.ExpiresAt
		return nil
	})
}
