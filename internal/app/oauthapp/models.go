package oauthapp

import (
	"net/url"
	"time"
)

type AuthorizationURLResponse struct {
	URL *url.URL
}

type OAuthApp struct {
	ID         string
	OwnerID    string
	ProviderID string
	Scopes     []string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
