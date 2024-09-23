package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Provider struct {
	Name         string   `yaml:"name"`
	ClientID     string   `yaml:"client_id"`
	ClientSecret string   `yaml:"client_secret"`
	RedirectURL  string   `yaml:"redirect_url"`
	AuthURL      string   `yaml:"auth_url"`
	TokenURL     string   `yaml:"token_url"`
	Scopes       []string `yaml:"scopes"`
}

type Config struct {
	Providers                 []Provider `yaml:"providers"`
	AllowProviderRegistration bool       `yaml:"allow_provider_registration"`
}

func NewConfig(cfgPath string) (*Config, error) {

	viper.SetConfigName("config")
	viper.AddConfigPath(cfgPath)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var configObject Config
	if err := viper.Unmarshal(&configObject); err != nil {
		return nil, err
	}

	return &configObject, nil
}
