package main

import (
	"context"
	"github.com/microserv-io/oauth-credentials-server/internal/config"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/provider"
	"github.com/microserv-io/oauth-credentials-server/internal/infrastructure/gorm"
	"github.com/microserv-io/oauth-credentials-server/internal/infrastructure/grpc"
	"github.com/microserv-io/oauth-credentials-server/internal/usecase"
	"log"
	"net"
	"os"
)

const CfgPath = "/cfg"

func main() {

	configObj, err := config.NewConfig(CfgPath)
	if err != nil {
		panic(err)
	}

	log.Printf("Configuration loaded: %+v", configObj)

	db, err := gorm.Open(os.Getenv("DATABASE_URL"), true)
	if err != nil {
		panic(err)
	}

	oauthAppRepository := gorm.NewOAuthAppRepository(db)
	providerRepository := gorm.NewProviderRepository(db)

	var providers []*provider.Provider
	for _, p := range configObj.Providers {
		providers = append(providers, &provider.Provider{
			Name:         p.Name,
			ClientID:     p.ClientID,
			ClientSecret: p.ClientSecret,
			RedirectURL:  p.RedirectURL,
			AuthURL:      p.AuthURL,
			TokenURL:     p.TokenURL,
			Scopes:       p.Scopes,
		})
	}

	if err := usecase.NewLoadProvidersUseCase(
		providerRepository,
	).Execute(
		context.Background(), providers,
	); err != nil {
		log.Fatalf("[STARTUP] failed to load providers: %v", err)
	}

	server := grpc.NewServer(
		usecase.NewListOAuthUseCase(oauthAppRepository),
		usecase.NewGetCredentialsUseCase(oauthAppRepository, nil),
	)

	log.Printf("Starting server on port 8080")

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
