package internal

import (
	"context"
	"fmt"
	"github.com/microserv-io/oauth-credentials-server/internal/app/oauthapp"
	"github.com/microserv-io/oauth-credentials-server/internal/app/provider"
	"github.com/microserv-io/oauth-credentials-server/internal/config"
	domainprovider "github.com/microserv-io/oauth-credentials-server/internal/domain/models/provider"
	"github.com/microserv-io/oauth-credentials-server/internal/infrastructure/encryption"
	"github.com/microserv-io/oauth-credentials-server/internal/infrastructure/gorm"
	grpcimpl "github.com/microserv-io/oauth-credentials-server/internal/infrastructure/grpc"
	"github.com/microserv-io/oauth-credentials-server/internal/infrastructure/oauth2"
	"github.com/microserv-io/oauth-credentials-server/internal/usecase"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"net"
	"os"
)

const defaultPortNumber = "8080"

type Application struct {
	server *grpc.Server
	config *config.Config
}

func NewApplication(cfgPath string) (*Application, error) {

	configObj, err := config.NewConfig(cfgPath, "config")
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %v", err)
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

	var providers []*domainprovider.Provider
	for _, p := range configObj.Providers {
		providers = append(providers, &domainprovider.Provider{
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
		return nil, fmt.Errorf("failed to load providers: %v", err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	encryptor, err := encryption.NewAesGcmEncryptor("random-key")
	if err != nil {
		return nil, fmt.Errorf("failed to create encryptor: %v", err)
	}

	oauthAppService := oauthapp.NewService(
		oauthAppRepository,
		providerRepository,
		&oauth2.TokenSourceFactory{},
		logger,
	)

	providerService := provider.NewService(
		providerRepository,
		oauthAppRepository,
		encryptor,
	)

	server := grpcimpl.NewServer(
		oauthAppService,
		providerService,
		logger,
	)

	return &Application{
		server: server,
		config: configObj,
	}, nil
}

func (a *Application) Run(portNumber string) error {

	if portNumber == "" {
		portNumber = "8080"
	}

	log.Printf("Starting server on port %s", portNumber)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", portNumber))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	if err := a.server.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}

func (a *Application) Stop() {
	a.server.GracefulStop()
}
