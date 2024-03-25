package core

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection() (*gorm.DB, error) {
	gormConfig := &gorm.Config{}
	dns := "host=db user=thinhlh password=thinhlh dbname=go_ecommerce port=5432"
	return gorm.Open(postgres.Open(dns), gormConfig)
}
