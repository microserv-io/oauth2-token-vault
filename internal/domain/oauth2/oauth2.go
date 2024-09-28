package oauth2

import (
	"context"
)

type Client interface {
	Exchange(ctx context.Context, config *Config, code string) (*Token, error)
}
