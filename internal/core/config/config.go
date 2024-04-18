package config

import "github.com/spf13/viper"

type Config struct {
	Database   DatabaseConfig
	ServerPort int `mapstructure:"SERVER_PORT"`
}

func newViper() *viper.Viper {
	v := viper.New()

	v.AutomaticEnv()
	v.AddConfigPath(".")
	v.SetConfigType("env")
	v.SetConfigFile(".env")

	err := v.ReadInConfig()

	if err != nil {
		panic(err)
	}

	return v
}

func LoadConfig() (Config, error) {
	v := newViper()
	var cfg Config

	databaseConfig, err := loadDatabaseConfig(v)
	if err != nil {
		return cfg, err
	}

	cfg = Config{
		Database: databaseConfig,
	}

	err = v.Unmarshal(&cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, err

}
