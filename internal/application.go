package internal

import (
	"context"
	"fmt"
	"github.com/microserv-io/oauth2-token-vault/internal/app/oauthapp"
	"github.com/microserv-io/oauth2-token-vault/internal/app/provider"
	"github.com/microserv-io/oauth2-token-vault/internal/config"
	domainprovider "github.com/microserv-io/oauth2-token-vault/internal/domain/models/provider"
	"github.com/microserv-io/oauth2-token-vault/internal/infrastructure/encryption"
	gormimpl "github.com/microserv-io/oauth2-token-vault/internal/infrastructure/gorm"
	grpcimpl "github.com/microserv-io/oauth2-token-vault/internal/infrastructure/grpc"
	"github.com/microserv-io/oauth2-token-vault/internal/infrastructure/oauth2"
	"github.com/microserv-io/oauth2-token-vault/internal/usecase"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log/slog"
	"net"
	"os"
)

type Application struct {
	server *grpc.Server
	lis    net.Listener
	config *config.Config
	db     *gorm.DB
}

type Option func(*Application) error

func WithDatabase(db *gorm.DB) Option {
	return func(a *Application) error {
		a.db = db
		return nil
	}
}

func NewApplication(cfgPath string, opts ...Option) (*Application, error) {
	app := &Application{}
	for _, option := range opts {
		if err := option(app); err != nil {
			return nil, fmt.Errorf("failed to apply option: %v", err)
		}
	}
	configObj, err := config.NewConfig(cfgPath, "config")
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %v", err)
	}
	app.config = configObj

	if app.db == nil {
		db, err := gormimpl.Open(
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
		app.db = db
	}

	oauthAppRepository := gormimpl.NewOAuthAppRepository(app.db)
	providerRepository := gormimpl.NewProviderRepository(app.db)

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
		oauth2.NewClient(),
	)

	app.server = grpcimpl.NewServer(
		oauthAppService,
		providerService,
		logger,
	)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configObj.Port))
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %v", err)
	}

	app.lis = lis

	return app, nil
}

func (a *Application) GetPort() int {
	return a.lis.Addr().(*net.TCPAddr).Port
}

// Run starts the gRPC server on the specified port, or 8080 if not specified.
func (a *Application) Run() error {
	if err := a.server.Serve(a.lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}

func (a *Application) Stop() {
	a.server.GracefulStop()
}
