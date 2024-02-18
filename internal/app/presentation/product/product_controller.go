package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	product "github.com/thinhlh/go-web-101/internal/app/application/product"
)

func FindAllProducts(c *gin.Context) {
	products := product.FindAllProducts()

	c.JSON(http.StatusAccepted, products)
}
