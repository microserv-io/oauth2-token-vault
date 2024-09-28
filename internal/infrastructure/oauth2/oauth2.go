package oauth2

import (
	"context"
	"fmt"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/oauth2"
	oauth "golang.org/x/oauth2"
)

type StdOauth2Client struct {
}

func NewStdOauth2Client() *StdOauth2Client {
	return &StdOauth2Client{}
}

func (c *StdOauth2Client) Exchange(ctx context.Context, config *oauth2.Config, code string) (*oauth2.Token, error) {

	client := oauth.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Endpoint: oauth.Endpoint{
			AuthURL:  config.AuthURL,
			TokenURL: config.TokenURL,
		},
		RedirectURL: config.RedirectURL,
		Scopes:      config.Scopes,
	}

	token, err := client.Exchange(ctx, code)

	if err != nil {
		return nil, fmt.Errorf("could not exchange code for token: %w", err)
	}

	return &oauth2.Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresAt:    token.Expiry,
	}, nil
}
