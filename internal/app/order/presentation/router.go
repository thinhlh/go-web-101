package presentation

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/thinhlh/go-web-101/internal/core/config"
	"gorm.io/gorm"
)

func NewOrderRouter(config config.Config, connection *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	orderController := NewOrderController()

	r.Route("/api/v1/orders", func(r chi.Router) {
		r.Get("/", orderController.Ping)
	})

	return r
}
