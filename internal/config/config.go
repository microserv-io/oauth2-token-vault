package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
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

type Config struct {
	Providers                 []Provider
	AllowProviderRegistration bool `mapstructure:"allow_provider_registration"`
}

func NewConfig(cfgPath string, configFileName string) (*Config, error) {

	if configFileName == "" {
		configFileName = "config"
	}

	viper.SetConfigName(configFileName)
	viper.AddConfigPath(cfgPath)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var configObject Config
	if err := viper.Unmarshal(&configObject); err != nil {
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
