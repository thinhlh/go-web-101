package main

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/thinhlh/go-web-101/internal/core/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	m, err := migrate.New(
		"./migrations",
		fmt.Sprintf(
			"postgres://%v:%v@%v:%v/%v",
			config.Database.DatabaseHost,
			config.Database.DatabasePort,
			config.Database.DatabaseUsername,
			config.Database.DatabasePassword,
			config.Database.DatabaseName,
		),
	)

	if err != nil {
		panic(err)
	}

	m.Version()
}
