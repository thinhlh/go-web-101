package config

import (
	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	DatabaseHost     string `mapstructure:"POSTGRES_HOST"`
	DatabasePort     int    `mapstructure:"POSTGRES_PORT"`
	DatabaseName     string `mapstructure:"POSTGRES_DB"`
	DatabaseUsername string `mapstructure:"POSTGRES_USER"`
	DatabasePassword string `mapstructure:"POSTGRES_PASSWORD"`
}

func newViperDatabaseConfig() *viper.Viper {
	v := viper.New()

	v.AddConfigPath(".")
	v.SetConfigType("env")
	v.SetConfigFile(".env")

	err := v.ReadInConfig()

	if err != nil {
		panic(err)
	}
	v.AutomaticEnv()

	return v
}

func loadDatabaseConfig() (DatabaseConfig, error) {
	v := newViperDatabaseConfig()

	var cfg DatabaseConfig
	err := v.Unmarshal(&cfg)

	return cfg, err
}
