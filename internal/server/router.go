package server

import (
	"github.com/gin-gonic/gin"
	"github.com/thinhlh/go-web-101/internal/app/product/presentation"
)

func NewRouter(productController presentation.ProductController) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

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
