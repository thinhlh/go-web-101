package product

import (
	domain "github.com/thinhlh/go-web-101/internal/app/domain/product"
	"github.com/thinhlh/go-web-101/internal/app/infrastructure/product"
)

func FindAllProducts() []domain.Product {
	r := product.NewRepository()

	return r.FindAllProducts()
}
