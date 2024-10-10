package oauthapp

import (
	"net/url"
	"time"
)

type CreateAuthorizationURLForProviderResponse struct {
	URL *url.URL
}

type RetrieveAccessTokenResponse struct {
	AccessToken string
}

type ListOAuthAppsForOwnerResponse struct {
	Apps []*OAuthApp
}

type GetOAuthForProviderAndOwnerResponse struct {
	App *OAuthApp
}

type OAuthApp struct {
	ID         string
	OwnerID    string
	ProviderID string
	Scopes     []string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
