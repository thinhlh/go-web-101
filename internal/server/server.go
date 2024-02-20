package server

import (
	"errors"
	"log"
	"net/http"

	"github.com/thinhlh/go-web-101/internal/app/product/application"
	"github.com/thinhlh/go-web-101/internal/app/product/infrastructure"
	"github.com/thinhlh/go-web-101/internal/app/product/presentation"
	"github.com/thinhlh/go-web-101/internal/core"
)

func New() {

	connection, err := core.NewDatabaseConnection()

	if err != nil {
		log.Fatalf("unable to connect to database, %v", err)
	}

	router := NewRouter(
		presentation.NewProductController(
			application.NewProductService(
				infrastructure.NewProductRepository(
					connection,
				),
			),
		),
	)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// go func() {
	// 	// service connections
	// 	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
	// 		log.Fatalf("listen: %s\n", err)
	// 	}
	// }()

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("listen: %s\n", err)
	}
}
