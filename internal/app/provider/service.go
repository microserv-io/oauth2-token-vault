package provider

import (
	"context"
	"fmt"
	"github.com/microserv-io/oauth-credentials-server/internal/domain"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/oauthapp"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/provider"
)

// Encryptor is an interface for encrypting and decrypting data.
type Encryptor interface {
	domain.Encryptor
}

// ProviderRepository provides access to the provider storage.
type ProviderRepository interface {
	provider.Repository
}

// OAuthAppRepository provides access to the oauth app storage.
type OAuthAppRepository interface {
	oauthapp.Repository
}

// Service provides provider operations.
type Service struct {
	providerRepository ProviderRepository
	oauthAppRepository OAuthAppRepository
	encryptor          Encryptor
}

// NewService creates a new provider service.
func NewService(
	providerRepository ProviderRepository,
	oauthAppRepository OAuthAppRepository,
	encryptor Encryptor,
) *Service {

	if providerRepository == nil {
		panic("providerRepository is required")
	}

	if oauthAppRepository == nil {
		panic("oauthAppRepository is required")
	}

	if encryptor == nil {
		panic("encryptor is required")
	}

	return &Service{
		providerRepository: providerRepository,
		oauthAppRepository: oauthAppRepository,
		encryptor:          encryptor,
	}
}

// GetProviderByName returns a provider by name.
func (s *Service) GetProviderByName(ctx context.Context, name string) (*GetProviderByNameResponse, error) {
	p, err := s.providerRepository.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return &GetProviderByNameResponse{
		Provider: &Provider{
			Name:        p.Name,
			AuthURL:     p.AuthURL,
			TokenURL:    p.TokenURL,
			Scopes:      p.Scopes,
			RedirectURI: p.RedirectURL,
			ClientID:    p.ClientID,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
		},
	}, nil
}

// ListProviders lists all providers.
func (s *Service) ListProviders(ctx context.Context) (*ListProvidersResponse, error) {
	providers, err := s.providerRepository.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list providers: %w", err)
	}

	r := &ListProvidersResponse{
		Providers: make([]*Provider, 0, len(providers)),
	}

	for _, p := range providers {
		r.Providers = append(r.Providers, &Provider{
			Name:        p.Name,
			AuthURL:     p.AuthURL,
			TokenURL:    p.TokenURL,
			Scopes:      p.Scopes,
			RedirectURI: p.RedirectURL,
			ClientID:    p.ClientID,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
		})
	}

	return r, nil

}

// CreateProvider creates a new provider.
func (s *Service) CreateProvider(ctx context.Context, input *CreateInput, source string) (*CreateProviderResponse, error) {
	providerObj, err := provider.NewProvider(input.Name, input.ClientID, input.ClientSecret, input.RedirectURI, input.AuthURL, input.TokenURL, input.Scopes, source)
	if err != nil {
		return nil, fmt.Errorf("failed to create provider: %w", err)
	}

	if err := s.providerRepository.Create(ctx, &providerObj); err != nil {
		return nil, fmt.Errorf("failed to create provider: %w", err)
	}

	return &CreateProviderResponse{
		Provider: &Provider{
			Name:        providerObj.Name,
			AuthURL:     providerObj.AuthURL,
			TokenURL:    providerObj.TokenURL,
			Scopes:      providerObj.Scopes,
			RedirectURI: providerObj.RedirectURL,
			ClientID:    providerObj.ClientID,
			CreatedAt:   providerObj.CreatedAt,
			UpdatedAt:   providerObj.UpdatedAt,
		},
	}, nil
}

// Update updates a provider by name.
func (s *Service) Update(ctx context.Context, name string, input *UpdateInput) (*UpdateProviderResponse, error) {
	providerObj, err := s.providerRepository.FindByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("failed to find provider by name: %w", err)
	}

	encryptedClientSecret, err := s.encryptor.Encrypt(input.ClientSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt client secret: %w", err)
	}

	providerObj.ClientID = input.ClientID
	providerObj.ClientSecret = encryptedClientSecret
	providerObj.RedirectURL = input.RedirectURI
	providerObj.AuthURL = input.AuthURL
	providerObj.TokenURL = input.TokenURL
	providerObj.Scopes = input.Scopes

	if err := s.providerRepository.Update(ctx, providerObj); err != nil {
		return nil, fmt.Errorf("failed to update provider: %w", err)
	}

	return &UpdateProviderResponse{
		Provider: &Provider{
			Name:        providerObj.Name,
			AuthURL:     providerObj.AuthURL,
			TokenURL:    providerObj.TokenURL,
			Scopes:      providerObj.Scopes,
			RedirectURI: providerObj.RedirectURL,
			ClientID:    providerObj.ClientID,
			CreatedAt:   providerObj.CreatedAt,
			UpdatedAt:   providerObj.UpdatedAt,
		},
	}, nil
}

// DeleteProvider deletes a provider by name. It returns an error if the provider has associated oauth apps.
func (s *Service) DeleteProvider(ctx context.Context, name string) error {
	apps, err := s.oauthAppRepository.ListForProvider(ctx, name)
	if err != nil {
		return fmt.Errorf("failed to list oauth apps for provider: %w", err)
	}

	if len(apps) > 0 {
		return fmt.Errorf("provider has associated oauth apps")
	}

	if err := s.providerRepository.Delete(ctx, name); err != nil {
		return fmt.Errorf("failed to delete provider: %w", err)
	}

	return nil
}