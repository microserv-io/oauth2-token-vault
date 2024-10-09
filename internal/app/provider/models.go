package provider

import "time"

type Provider struct {
	ID          string
	Name        string
	ClientID    string
	Scopes      []string
	RedirectURI string
	AuthURL     string
	TokenURL    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CreateInput struct {
	Name         string
	ClientID     string
	ClientSecret string
	RedirectURI  string
	AuthURL      string
	TokenURL     string
	Scopes       []string
}

type UpdateInput struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	AuthURL      string
	TokenURL     string
	Scopes       []string
}

type CreateProviderResponse struct {
	Provider *Provider
}

type UpdateProviderResponse struct {
	Provider *Provider
}

type ListProvidersResponse struct {
	Providers []*Provider
}

type GetProviderByNameResponse struct {
	Provider *Provider
}
