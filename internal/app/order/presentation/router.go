package presentation

import (
	"github.com/go-chi/chi/v5"
	"github.com/thinhlh/go-web-101/internal/core/config"
	"github.com/thinhlh/go-web-101/internal/core/middlewares"
	"gorm.io/gorm"
)

func NewOrderRouter(config config.Config, connection *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	middlewares.AttachRestMiddlewares(r, config)

	orderController := NewOrderController()

	r.Route("/api/v1/orders", func(r chi.Router) {
		r.Get("/", orderController.Ping)
	})

	return r
}
