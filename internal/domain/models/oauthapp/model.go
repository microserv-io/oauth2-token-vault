package oauthapp

import "time"

type OAuthApp struct {
	ID           string
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
