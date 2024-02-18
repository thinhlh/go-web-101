package server

import (
	"github.com/gin-gonic/gin"
	"github.com/thinhlh/go-web-101/internal/app/presentation/category"
	"github.com/thinhlh/go-web-101/internal/app/presentation/product"
)

func New() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			products := v1.Group("/products")
			{
				products.GET("/", product.FindAllProducts)
			}

			categories := v1.Group("/categories")
			{
				categories.GET("/", category.FindAllCategories)
			}
		}
	}

	r.Run(":8080")
}
