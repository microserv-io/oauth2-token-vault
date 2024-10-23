package provider

import (
	"context"
	"fmt"
	"github.com/microserv-io/oauth2-token-vault/internal/domain"
	"github.com/microserv-io/oauth2-token-vault/internal/domain/models/oauthapp"
	"github.com/microserv-io/oauth2-token-vault/internal/domain/models/provider"
	"github.com/microserv-io/oauth2-token-vault/internal/domain/oauth2"
	"log/slog"
	"net/url"
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

// OAuth2Client is an interface for exchanging authorization codes for tokens.
type OAuth2Client interface {
	oauth2.Client
}

// Service provides provider operations.
type Service struct {
	providerRepository ProviderRepository
	oauthAppRepository OAuthAppRepository
	encryptor          Encryptor
	oauth2Client       OAuth2Client
}

// NewService creates a new provider service.
func NewService(
	providerRepository ProviderRepository,
	oauthAppRepository OAuthAppRepository,
	encryptor Encryptor,
	oauth2Client OAuth2Client,
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

	if oauth2Client == nil {
		panic("oauth2Client is required")
	}

	return &Service{
		providerRepository: providerRepository,
		oauthAppRepository: oauthAppRepository,
		encryptor:          encryptor,
		oauth2Client:       oauth2Client,
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

// SyncProviders syncs a list of providers with the repository. Providers not in the list are deleted.
func (s *Service) SyncProviders(ctx context.Context, request *SyncProviderRequest) error {
	existingProviders, err := s.providerRepository.List(ctx)
	if err != nil {
		return fmt.Errorf("failed to list existing providers: %w", err)
	}

	providersMap := make(map[string]*SyncProvider)
	for _, p := range request.Providers {
		providersMap[p.Name] = p
	}

	for _, existingProvider := range existingProviders {

		if existingProvider.Source != provider.SourceConfig {
			continue
		}

		if _, found := providersMap[existingProvider.Name]; !found {
			if err := s.providerRepository.Delete(ctx, existingProvider.Name); err != nil {
				slog.Warn("failed to delete provider")
			}
		}
	}

	for _, providerToSync := range request.Providers {

		clientSecret, err := s.encryptor.Encrypt(providerToSync.ClientSecret)
		if err != nil {
			return fmt.Errorf("failed to encrypt client secret: %w", err)
		}
		providerObj, err := provider.NewProvider(
			providerToSync.Name,
			providerToSync.ClientID,
			clientSecret, providerToSync.RedirectURI,
			providerToSync.AuthURL,
			providerToSync.TokenURL,
			providerToSync.Scopes,
			"config",
		)
		if err != nil {
			return fmt.Errorf("failed to instantiate provider: %w", err)
		}
		if _, err := s.providerRepository.FindByName(ctx, providerObj.Name); err != nil {
			if err := s.providerRepository.Create(ctx, providerObj); err != nil {
				return err
			}
		} else {
			if err := s.providerRepository.Update(ctx, providerObj); err != nil {
				return err
			}
		}
	}

	return nil
}

// CreateProvider creates a new provider.
func (s *Service) CreateProvider(ctx context.Context, request *CreateProviderRequest, source string) (*CreateProviderResponse, error) {
	providerObj, err := provider.NewProvider(request.Name, request.ClientID, request.ClientSecret, request.RedirectURI, request.AuthURL, request.TokenURL, request.Scopes, source)
	if err != nil {
		return nil, fmt.Errorf("failed to create provider: %w", err)
	}

	if err := s.providerRepository.Create(ctx, providerObj); err != nil {
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

// UpdateProvider updates a provider by name.
func (s *Service) UpdateProvider(ctx context.Context, name string, request *UpdateProviderRequest) (*UpdateProviderResponse, error) {
	providerObj, err := s.providerRepository.FindByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("failed to find provider by name: %w", err)
	}

	encryptedClientSecret, err := s.encryptor.Encrypt(request.ClientSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt client secret: %w", err)
	}

	providerObj.ClientID = request.ClientID
	providerObj.ClientSecret = encryptedClientSecret
	providerObj.RedirectURL = request.RedirectURI
	providerObj.AuthURL = request.AuthURL
	providerObj.TokenURL = request.TokenURL
	providerObj.Scopes = request.Scopes

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
func (s *Service) DeleteProvider(ctx context.Context, request *DeleteProviderRequest) error {
	apps, err := s.oauthAppRepository.ListForProvider(ctx, request.Name)
	if err != nil {
		return fmt.Errorf("failed to list oauth apps for provider: %w", err)
	}

	if len(apps) > 0 && !request.DeleteConnectedOAuthApps {
		return fmt.Errorf("provider has associated oauth apps")
	}

	for _, app := range apps {
		if err := s.oauthAppRepository.Delete(ctx, app.ID); err != nil {
			return fmt.Errorf("failed to delete oauth app: %w", err)
		}
	}

	if err := s.providerRepository.Delete(ctx, request.Name); err != nil {
		return fmt.Errorf("failed to delete provider: %w", err)
	}

	return nil
}

// GetAuthorizationURL returns the authorization URL for a provider.
func (s *Service) GetAuthorizationURL(ctx context.Context, input *GetAuthorizationURLRequest) (*GetAuthorizationURLResponse, error) {
	providerObj, err := s.providerRepository.FindByName(ctx, input.Provider)
	if err != nil {
		return nil, fmt.Errorf("failed to find provider by name: %w", err)
	}

	oauth2Config := &oauth2.Config{
		ClientID:     providerObj.ClientID,
		ClientSecret: providerObj.ClientSecret,
		RedirectURL:  providerObj.RedirectURL,
		Scopes:       providerObj.Scopes,
		AuthURL:      providerObj.AuthURL,
		TokenURL:     providerObj.TokenURL,
	}

	redirectURL, err := s.oauth2Client.GetAuthorizationURL(oauth2Config, input.State)
	if err != nil {
		return nil, fmt.Errorf("failed to get authorization url: %w", err)
	}
	if redirectURL == "" {
		return nil, fmt.Errorf("failed to create redirect url for provider: %s", input.Provider)
	}

	parsedURL, err := url.Parse(redirectURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse redirect url: %w", err)
	}

	return &GetAuthorizationURLResponse{
		URL: parsedURL,
	}, nil
}

// ExchangeAuthorizationCode exchanges an authorization code for an access token.
func (s *Service) ExchangeAuthorizationCode(ctx context.Context, input *ExchangeAuthorizationCodeRequest) error {
	providerObj, err := s.providerRepository.FindByName(ctx, input.Provider)
	if err != nil {
		return fmt.Errorf("failed to find providerObj by name: %w", err)
	}

	token, err := s.oauth2Client.Exchange(ctx, &oauth2.Config{
		ClientID:     providerObj.ClientID,
		ClientSecret: providerObj.ClientSecret,
		AuthURL:      providerObj.AuthURL,
		TokenURL:     providerObj.TokenURL,
		RedirectURL:  providerObj.RedirectURL,
		Scopes:       providerObj.Scopes,
	}, input.Code)
	if err != nil {
		return fmt.Errorf("failed to exchange authorization code: %w", err)
	}

	accessToken, err := s.encryptor.Encrypt(token.AccessToken)
	if err != nil {
		return fmt.Errorf("failed to encrypt access token: %w", err)
	}

	refreshToken, err := s.encryptor.Encrypt(token.RefreshToken)
	if err != nil {
		return fmt.Errorf("failed to encrypt refresh token: %w", err)
	}

	oauthApp := oauthapp.NewOAuthApp(
		providerObj.Name,
		accessToken,
		refreshToken,
		token.TokenType,
		token.ExpiresAt,
		providerObj.Scopes,
		input.OwnerID,
	)

	if err := s.oauthAppRepository.Create(ctx, oauthApp); err != nil {
		return fmt.Errorf("failed to create oauth app: %w", err)
	}

	return nil
}
