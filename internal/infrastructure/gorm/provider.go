package gorm

import (
	"context"
	"fmt"
	"github.com/lib/pq"
	"github.com/microserv-io/oauth2-token-vault/internal/domain/models/provider"
	"gorm.io/gorm"
	"time"
)

type Provider struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	Name         string `gorm:"unique"`
	ClientID     string
	ClientSecret string
	RedirectURL  string
	AuthURL      string
	TokenURL     string
	Scopes       pq.StringArray `gorm:"type:text[]"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Provider) TableName() string {
	return "providers"
}

func (dao Provider) ToDomain() *provider.Provider {
	return &provider.Provider{
		ID:           dao.ID,
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
		ID:           provider.ID,
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

func NewProviderRepository(db *gorm.DB) *ProviderRepository {
	return &ProviderRepository{
		db: db,
	}
}

func (p ProviderRepository) FindByName(ctx context.Context, name string) (*provider.Provider, error) {
	var providerDao Provider
	if err := p.db.WithContext(ctx).Where("name = ?", name).First(&providerDao).Error; err != nil {
		return nil, err
	}

	return providerDao.ToDomain(), nil
}

func (p ProviderRepository) List(ctx context.Context) ([]*provider.Provider, error) {
	providerDaos := make([]Provider, 0)

	if err := p.db.WithContext(ctx).Find(&providerDaos).Error; err != nil {
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

	if err := p.db.WithContext(ctx).Create(providerDao).Error; err != nil {
		return err
	}

	return nil
}

func (p ProviderRepository) Delete(ctx context.Context, name string) error {
	q := p.db.WithContext(ctx).Where("name = ?", name).Delete(&Provider{})

	if q.Error != nil {
		return fmt.Errorf("failed to delete provider: %w", q.Error)
	}

	if q.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (p ProviderRepository) Update(ctx context.Context, provider *provider.Provider) error {
	providerDao := newProviderFromDomain(provider)

	if err := p.db.WithContext(ctx).Save(providerDao).Error; err != nil {
		return err
	}

	return nil
}
