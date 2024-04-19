package presentation

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/thinhlh/go-web-101/internal/app/product/application"
	"github.com/thinhlh/go-web-101/internal/app/product/infrastructure"
	"github.com/thinhlh/go-web-101/internal/core/config"
	"gorm.io/gorm"
)

func NewProductRouter(config config.Config, connection *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	productController := NewProductController(
		application.NewProductService(
			infrastructure.NewProductRepository(
				connection,
			),
		),
	)

	r.Route("/api/v1/products", func(r chi.Router) {
		r.Get("/", productController.GetProducts)
		r.Get("/{id}", productController.GetProductById)
	})

	return r
}
