package oauthapp

import "context"

type Repository interface {
	Find(ctx context.Context, ownerID string, id string) (*OAuthApp, error)
	ListForOwner(ctx context.Context, ownerID string) ([]*OAuthApp, error)
	Create(ctx context.Context, app *OAuthApp) error
	UpdateByID(ctx context.Context, id string, updateFn func(app *OAuthApp) error) error
	Delete(ctx context.Context, id string) error
}
