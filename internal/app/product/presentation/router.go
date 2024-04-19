package presentation

import (
	"github.com/go-chi/chi/v5"
	"github.com/thinhlh/go-web-101/internal/app/product/application"
	"github.com/thinhlh/go-web-101/internal/app/product/infrastructure"
	"github.com/thinhlh/go-web-101/internal/core/config"
	"github.com/thinhlh/go-web-101/internal/core/database"
	"github.com/thinhlh/go-web-101/internal/core/middlewares"
)

func NewProductRouter(config config.Config, connection *database.Database) *chi.Mux {
	r := chi.NewRouter()

	middlewares.AttachRestMiddlewares(r, config)

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
