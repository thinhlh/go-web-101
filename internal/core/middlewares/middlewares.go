package middlewares

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/thinhlh/go-web-101/internal/core/config"
)

func AttachRestMiddlewares(router *chi.Mux, config config.Config) {
	router.Use(middleware.Throttle(15))
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	// router.Use(middleware.CleanPath)
	router.Use(middleware.Heartbeat("/"))

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))
	router.Use(httprate.LimitByIP(config.ApiLimitByIpPerMin, 1*time.Minute))

	router.Mount("/profiler", middleware.Profiler())
}
