package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Name string `envconfig:"NAME" default:""`

	Port string `envconfig:"PORT" default:""`

	DBHost string `envconfig:"DB_HOST" default:""`
	DBPort string `envconfig:"DB_PORT" default:""`
	DBName string `envconfig:"DB_NAME" default:""`
	DBUser string `envconfig:"DB_USER" default:""`
	DBPass string `envconfig:"DB_PASS" default:""`

	NewRelicLicenseKey string `envconfig:"NEW_RELIC_LICENSE_KEY" default:""`
}

func (c Config) Host() string {
	return ":" + c.Port
}

func Load() (Config, error) {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		return config, err
	}

	return config, nil
}

func MustLoad() Config {
	var config Config
	envconfig.MustProcess("", &config)
	return config
}
