package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/thinhlh/go-web-101/internal/app/order/presentation"
	bootmanager "github.com/thinhlh/go-web-101/internal/core/boot_manager"
	"github.com/thinhlh/go-web-101/internal/core/config"
	"github.com/thinhlh/go-web-101/internal/core/database"
)

func New() bootmanager.Daemon {

	config, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("unable to load config, %v", err)
		panic(err)
	}

	connection, err := database.NewDatabaseConnection(config.Database)

	if err != nil {
		log.Fatalf("unable to connect to database, %v", err)
		panic(err)
	}

	router := presentation.NewOrderRouter(config, connection)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.OrderService.Port),
		Handler: router,
	}

	go func() {
		// Run server in another goroutine & let the returned function as clean up func
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Printf("server is listening on port %v", config.OrderService.Port)

	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// DO SHUTDOWN & CLEANUP
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server shutdown with error:", err)
		}

		db, err := connection.DB.DB()
		if err != nil || db.Close() != nil {
			log.Fatal("Cannot close DB connection")
		}

		select {
		// catching ctx.Done(). timeout of 5 seconds.
		case <-ctx.Done():
			log.Printf("Server exiting timeout!")
		default:
		}

		log.Printf("Server exit gracefully!")
	}
}
