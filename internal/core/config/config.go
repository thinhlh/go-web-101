package config

type Config struct {
	Database DatabaseConfig
}

func LoadConfig() (Config, error) {
	databaseConfig, err := loadDatabaseConfig()

	cfg := Config{
		Database: databaseConfig,
	}

	return cfg, err

}
