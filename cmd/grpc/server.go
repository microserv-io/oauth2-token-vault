package main

import (
	"context"
	"fmt"
	"github.com/microserv-io/oauth-credentials-server/internal/app/oauthapp"
	"github.com/microserv-io/oauth-credentials-server/internal/config"
	"github.com/microserv-io/oauth-credentials-server/internal/domain/models/provider"
	"github.com/microserv-io/oauth-credentials-server/internal/infrastructure/gorm"
	"github.com/microserv-io/oauth-credentials-server/internal/infrastructure/grpc"
	"github.com/microserv-io/oauth-credentials-server/internal/infrastructure/oauth2"
	"github.com/microserv-io/oauth-credentials-server/internal/usecase"
	"log"
	"log/slog"
	"net"
	"os"
)

const CfgPath = "/cfg"

func main() {

	configObj, err := config.NewConfig(CfgPath, "config")
	if err != nil {
		panic(err)
	}

	log.Printf("Configuration loaded from file.")

	db, err := gorm.Open(
		fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			configObj.Database.Host,
			configObj.Database.User,
			configObj.Database.Password,
			configObj.Database.Name,
			configObj.Database.Port,
		),
		true,
		6,
		5,
		nil,
	)
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

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	oauthAppService := oauthapp.NewService(
		oauthAppRepository,
		providerRepository,
		&oauth2.TokenSourceFactory{},
		logger,
	)

	server := grpc.NewServer(
		oauthAppService,
		logger,
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
