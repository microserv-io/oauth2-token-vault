package usecase

import (
	"context"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/oauthapp"
)

type ListOAuthUseCase struct {
	repository oauthapp.Repository
}

func NewListOAuthUseCase(repository oauthapp.Repository) *ListOAuthUseCase {
	return &ListOAuthUseCase{
		repository: repository,
	}
}

func (u *ListOAuthUseCase) Execute(ctx context.Context, ownerID string) ([]*oauthapp.OAuthApp, error) {
	oauthApps, err := u.repository.ListForOwner(ctx, ownerID)

	if err != nil {
		return nil, err
	}

	return oauthApps, nil
}
