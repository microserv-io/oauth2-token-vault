package oauthapp

import "time"

type OAuthApp struct {
	ID           uint
	Provider     string
	AccessToken  string
	RefreshToken string
	TokenType    string
	ExpiresAt    time.Time
	Scopes       []string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	OwnerID      string
}

func NewOAuthApp(provider, accessToken, refreshToken, tokenType string, expiresAt time.Time, scopes []string, ownerID string) *OAuthApp {
	return &OAuthApp{
		Provider:     provider,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    tokenType,
		ExpiresAt:    expiresAt,
		Scopes:       scopes,
		OwnerID:      ownerID,
	}
}
