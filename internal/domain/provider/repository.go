package provider

import "context"

type Repository interface {
	FindByName(ctx context.Context, name string) (*Provider, error)
	List(ctx context.Context) ([]*Provider, error)
	Create(ctx context.Context, provider *Provider) error
}
