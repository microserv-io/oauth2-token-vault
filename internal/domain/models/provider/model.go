package provider

import (
	"errors"
	"time"
)

type Source string

var (
	ErrInvalidSource = errors.New("invalid source")
)

const (
	SourceConfig Source = "config"
	SourceAPI    Source = "api"
)

type Provider struct {
	ID           uint
	Name         string
	ClientID     string
	ClientSecret string
	RedirectURL  string
	AuthURL      string
	Source       Source
	TokenURL     string
	Scopes       []string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewProvider(name, clientID, clientSecret, redirectURL, authURL, tokenURL string, scopes []string, source string) (*Provider, error) {
	if source != string(SourceConfig) && source != string(SourceAPI) {
		return nil, ErrInvalidSource
	}

	return &Provider{
		Name:         name,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		AuthURL:      authURL,
		Source:       Source(source),
		TokenURL:     tokenURL,
		Scopes:       scopes,
	}, nil
}
