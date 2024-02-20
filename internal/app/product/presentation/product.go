package presentation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinhlh/go-web-101/internal/app/product/application"
)

type ProductController struct {
	productService application.ProductService
}

func NewProductController(service application.ProductServiceImpl) ProductController {
	return ProductController{service}
}

func (controller ProductController) GetProducts(c *gin.Context) {
	result := controller.productService.GetAllProducts()

	c.JSON(http.StatusOK, result)
}

func (controller ProductController) GetProductById(c *gin.Context) {
	result := controller.productService.GetProductById()

	c.JSON(http.StatusOK, result)
}
