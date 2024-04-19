package config

import "github.com/spf13/viper"

type ProductServiceConfig struct {
	Host string `mapstructure:"PRODUCT_SERVICE_HOST"`
	Port int    `mapstructure:"PRODUCT_SERVICE_PORT"`
}

type OrderServiceConfig struct {
	Host string `mapstructure:"ORDER_SERVICE_HOST"`
	Port int    `mapstructure:"ORDER_SERVICE_PORT"`
}

func loadProductServiceConfig(v *viper.Viper) (ProductServiceConfig, error) {
	var cfg ProductServiceConfig
	err := v.Unmarshal(&cfg)

	return cfg, err
}

func loadOrderServiceConfig(v *viper.Viper) (OrderServiceConfig, error) {
	var cfg OrderServiceConfig
	err := v.Unmarshal(&cfg)

	return cfg, err
}
