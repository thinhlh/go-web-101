package config

import "github.com/spf13/viper"

type DatabaseConfig struct {
	DatabaseHost     string `mapstructure:"POSTGRES_HOST"`
	DatabasePort     int    `mapstructure:"POSTGRES_PORT"`
	DatabaseName     string `mapstructure:"POSTGRES_DB"`
	DatabaseUsername string `mapstructure:"POSTGRES_USER"`
	DatabasePassword string `mapstructure:"POSTGRES_PASSWORD"`
}

func loadDatabaseConfig(v *viper.Viper) (DatabaseConfig, error) {
	var cfg DatabaseConfig
	err := v.Unmarshal(&cfg)

	return cfg, err
}
