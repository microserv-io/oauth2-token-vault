package gorm

import (
	"context"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/provider"
	"gorm.io/gorm"
	"time"
)

type Provider struct {
	gorm.Model
	Name         string
	ClientID     string
	ClientSecret string
	RedirectURL  string
	AuthURL      string
	TokenURL     string
	Scopes       []string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (dao Provider) ToDomain() *provider.Provider {
	return &provider.Provider{
		Name:         dao.Name,
		ClientID:     dao.ClientID,
		ClientSecret: dao.ClientSecret,
		RedirectURL:  dao.RedirectURL,
		AuthURL:      dao.AuthURL,
		TokenURL:     dao.TokenURL,
		Scopes:       dao.Scopes,
		CreatedAt:    dao.CreatedAt,
		UpdatedAt:    dao.UpdatedAt,
	}
}

func newProviderFromDomain(provider *provider.Provider) *Provider {
	return &Provider{
		Name:         provider.Name,
		ClientID:     provider.ClientID,
		ClientSecret: provider.ClientSecret,
		RedirectURL:  provider.RedirectURL,
		AuthURL:      provider.AuthURL,
		TokenURL:     provider.TokenURL,
		Scopes:       provider.Scopes,
		CreatedAt:    provider.CreatedAt,
		UpdatedAt:    provider.UpdatedAt,
	}
}

type ProviderRepository struct {
	db *gorm.DB
}

func (p ProviderRepository) FindByName(ctx context.Context, name string) (*provider.Provider, error) {
	var providerDao Provider
	if err := p.db.Where("name = ?", name).First(&providerDao).Error; err != nil {
		return nil, err
	}

	return providerDao.ToDomain(), nil
}

func (p ProviderRepository) List(ctx context.Context) ([]*provider.Provider, error) {
	providerDaos := make([]Provider, 0)

	if err := p.db.Find(&providerDaos).Error; err != nil {
		return nil, err
	}

	var result []*provider.Provider
	for _, providerDao := range providerDaos {
		result = append(result, providerDao.ToDomain())
	}

	return result, nil
}

func (p ProviderRepository) Create(ctx context.Context, provider *provider.Provider) error {
	providerDao := newProviderFromDomain(provider)

	if err := p.db.Create(providerDao).Error; err != nil {
		return err
	}

	return nil
}

func NewProviderRepository(db *gorm.DB) *ProviderRepository {
	return &ProviderRepository{
		db: db,
	}
}
