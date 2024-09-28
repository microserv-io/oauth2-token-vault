package usecase

import (
	"context"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/provider"
)

type LoadProvidersUseCase struct {
	providerRepository provider.Repository
}

func NewLoadProvidersUseCase(providerRepository provider.Repository) *LoadProvidersUseCase {
	return &LoadProvidersUseCase{
		providerRepository: providerRepository,
	}
}

func (u *LoadProvidersUseCase) Execute(ctx context.Context, providers []*provider.Provider) error {
	for _, providerObj := range providers {
		if _, err := u.providerRepository.FindByName(ctx, providerObj.Name); err != nil {
			if err := u.providerRepository.Create(ctx, providerObj); err != nil {
				return err
			}
		}
	}

	return nil
}
