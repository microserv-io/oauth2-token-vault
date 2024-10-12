package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type Provider struct {
	Name         string   `mapstructure:"name"`
	ClientID     string   `mapstructure:"client_id"`
	ClientSecret string   `mapstructure:"client_secret"`
	RedirectURL  string   `mapstructure:"redirect_url"`
	AuthURL      string   `mapstructure:"auth_url"`
	TokenURL     string   `mapstructure:"token_url"`
	Scopes       []string `mapstructure:"scopes"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

type Config struct {
	Port                      string     `mapstructure:"port"`
	Database                  Database   `mapstructure:"database"`
	Providers                 []Provider `mapstructure:"providers"`
	AllowProviderRegistration bool       `mapstructure:"allow_provider_registration"`
}

func NewConfig(cfgPath string, configFileName string) (*Config, error) {
	if configFileName == "" {
		configFileName = "config"
	}

	v := viper.NewWithOptions(viper.EnvKeyReplacer(strings.NewReplacer(".", "_")))

	v.SetConfigName(configFileName)
	v.AddConfigPath(cfgPath)
	v.AddConfigPath(".")

	v.AutomaticEnv()

	if err := v.BindEnv("port", "PORT"); err != nil {
		return nil, fmt.Errorf("error binding env variable: %w", err)
	}

	if err := v.BindEnv("database.host", "DATABASE_HOST"); err != nil {
		return nil, fmt.Errorf("error binding env variable: %w", err)
	}

	if err := v.BindEnv("database.port", "DATABASE_PORT"); err != nil {
		return nil, fmt.Errorf("error binding env variable: %w", err)
	}

	if err := v.BindEnv("database.user", "DATABASE_USER"); err != nil {
		return nil, fmt.Errorf("error binding env variable: %w", err)
	}

	if err := v.BindEnv("database.password", "DATABASE_PASSWORD"); err != nil {
		return nil, fmt.Errorf("error binding env variable: %w", err)
	}

	if err := v.BindEnv("database.name", "DATABASE_NAME"); err != nil {
		return nil, fmt.Errorf("error binding env variable: %w", err)
	}

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var configObject Config
	if err := v.Unmarshal(&configObject); err != nil {
		return nil, err
	}

	for i, provider := range configObject.Providers {
		if provider.Name == "" {
			return nil, fmt.Errorf("provider name is required for provider at index %d", i)
		}
		if provider.ClientID == "" {
			return nil, fmt.Errorf("client id is required for provider %s", provider.Name)
		}

		clientSecret := os.Getenv(fmt.Sprintf("PROVIDER__%d__CLIENT_SECRET", i))
		if clientSecret != "" {
			provider.ClientSecret = clientSecret
		}

		if provider.ClientSecret == "" {
			return nil, fmt.Errorf("client secret is required for provider %s", provider.Name)
		}

		if provider.RedirectURL == "" {
			return nil, fmt.Errorf("redirect url is required for provider %s", provider.Name)
		}
		if provider.AuthURL == "" {
			return nil, fmt.Errorf("auth url is required for provider %s", provider.Name)
		}
		if provider.TokenURL == "" {
			return nil, fmt.Errorf("token url is required for provider %s", provider.Name)
		}
		if len(provider.Scopes) == 0 {
			return nil, fmt.Errorf("scopes are required for provider %s", provider.Name)
		}

		configObject.Providers[i] = provider
	}

	return &configObject, nil
}
