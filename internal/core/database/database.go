package database

import (
	"fmt"

	"github.com/thinhlh/go-web-101/internal/core/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection(config config.DatabaseConfig) (*gorm.DB, error) {
	gormConfig := &gorm.Config{}

	dns := fmt.Sprintf(
		"host=%v user=%v  password=%v dbname=%v port=%v",
		config.DatabaseHost,
		config.DatabaseUsername,
		config.DatabasePassword,
		config.DatabaseName,
		config.DatabasePort,
	)
	return gorm.Open(postgres.Open(dns), gormConfig)
}
