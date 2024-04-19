package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/thinhlh/go-web-101/internal/app/product/application"
	"github.com/thinhlh/go-web-101/internal/core/dto"
)

type ProductController struct {
	productService application.ProductService
}

func NewProductController(service application.ProductService) ProductController {
	return ProductController{service}
}

func (controller ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	result := controller.productService.GetAllProducts()

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(dto.NewSuccessResponse(result))
}

func (controller ProductController) GetProductById(w http.ResponseWriter, r *http.Request) {
	result := controller.productService.GetProductById()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
