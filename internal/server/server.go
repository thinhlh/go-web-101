package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

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

	router := NewRouter(config, connection)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.ServerPort),
		Handler: router,
	}

	go func() {
		// Run server in another goroutine & let the returned function as clean up func
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// DO SHUTDOWN & CLEANUP
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server shutdown with error:", err)
		}

		select {
		// catching ctx.Done(). timeout of 5 seconds.
		case <-ctx.Done():
			log.Printf("Server exiting timeout!")
		default:
		}

		log.Printf("Server exit!")
	}
}
