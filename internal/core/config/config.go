package config

import "github.com/spf13/viper"

type Config struct {
	Database           DatabaseConfig
	ProductService     ProductServiceConfig
	OrderService       OrderServiceConfig
	ApiLimitByIpPerMin int `mapstructure:"API_LIMIT_BY_IP_PER_MIN"`
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
		panic(err)
	}

	productServiceConfig, err := loadProductServiceConfig(v)
	if err != nil {
		panic(err)
	}

	orderServiceConfig, err := loadOrderServiceConfig(v)
	if err != nil {
		panic(err)
	}

	cfg = Config{
		Database:       databaseConfig,
		ProductService: productServiceConfig,
		OrderService:   orderServiceConfig,
	}

	err = v.Unmarshal(&cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, err
}
