package server

import (
	"github.com/gin-gonic/gin"
	"github.com/thinhlh/go-web-101/internal/app/product/application"
	"github.com/thinhlh/go-web-101/internal/app/product/infrastructure"
	"github.com/thinhlh/go-web-101/internal/app/product/presentation"
	"github.com/thinhlh/go-web-101/internal/core/config"
	"gorm.io/gorm"
)

func NewRouter(config config.Config, connection *gorm.DB) *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	productController := presentation.NewProductController(
		application.NewProductService(
			infrastructure.NewProductRepository(
				connection,
			),
		),
	)

	api := r.Group("/api")
	{
		products := api.Group("/v1/products")
		{
			products.GET("/", productController.GetProducts)
			products.GET("/{id}", productController.GetProductById)
		}
	}

	return r
}
