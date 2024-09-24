package gorm

import (
	"github.com/lib/pq"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/oauthapp"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"time"
)

type OAuthApp struct {
	gorm.Model
	ID           string `gorm:"primaryKey"`
	Provider     string `gorm:"index"`
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
	Scopes       pq.StringArray `gorm:"type:text[]"`
	OwnerID      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (dao OAuthApp) ToDomain() *oauthapp.OAuthApp {

	return &oauthapp.OAuthApp{
		ID:           dao.ID,
		Provider:     dao.Provider,
		AccessToken:  dao.AccessToken,
		RefreshToken: dao.RefreshToken,
		ExpiresAt:    dao.ExpiresAt,
		Scopes:       dao.Scopes,
		OwnerID:      dao.OwnerID,
		CreatedAt:    dao.CreatedAt,
		UpdatedAt:    dao.UpdatedAt,
	}
}

func newOAuthAppFromDomain(app *oauthapp.OAuthApp) *OAuthApp {
	return &OAuthApp{
		ID:           app.ID,
		Provider:     app.Provider,
		AccessToken:  app.AccessToken,
		RefreshToken: app.RefreshToken,
		ExpiresAt:    app.ExpiresAt,
		Scopes:       app.Scopes,
		OwnerID:      app.OwnerID,
		CreatedAt:    app.CreatedAt,
		UpdatedAt:    app.UpdatedAt,
	}
}

type OAuthAppRepository struct {
	db *gorm.DB
}

func NewOAuthAppRepository(db *gorm.DB) *OAuthAppRepository {
	return &OAuthAppRepository{
		db: db,
	}
}

func (r OAuthAppRepository) Find(ctx context.Context, ownerID string, id string) (*oauthapp.OAuthApp, error) {

	var app OAuthApp

	if err := r.db.WithContext(ctx).Find(&app, "id = ? AND owner_id = ?", id, ownerID).Error; err != nil {
		return nil, err
	}

	return app.ToDomain(), nil
}

func (r OAuthAppRepository) ListForOwner(ctx context.Context, ownerID string) ([]*oauthapp.OAuthApp, error) {

	var apps []*OAuthApp

	if err := r.db.WithContext(ctx).Find(&apps, "owner_id = ?", ownerID).Error; err != nil {
		return nil, err
	}

	var result []*oauthapp.OAuthApp
	for _, app := range apps {
		result = append(result, app.ToDomain())
	}

	return result, nil
}

func (r OAuthAppRepository) Create(ctx context.Context, app *oauthapp.OAuthApp) error {
	if err := r.db.WithContext(ctx).Create(newOAuthAppFromDomain(app)).Error; err != nil {
		return err
	}

	return nil
}

func (r OAuthAppRepository) Update(ctx context.Context, app *oauthapp.OAuthApp) error {
	if err := r.db.WithContext(ctx).Save(newOAuthAppFromDomain(app)).Error; err != nil {
		return err
	}

	return nil
}

func (r OAuthAppRepository) Delete(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Delete(&OAuthApp{}, id).Error; err != nil {
		return err
	}

	return nil
}
