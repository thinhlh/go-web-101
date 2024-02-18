package core

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func InitDatabase() {
	gormConfig := &gorm.Config{}
	db, err := gorm.Open(postgres.Open("host=localhost user=thinhlh password=thinhlh dbname=go_ecommerce port=5432"), gormConfig)

	if err != nil {
		panic("Error when creating database connection")
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, _ := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Minute)

	Database = db
}
