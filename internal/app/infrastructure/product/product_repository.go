package product

import (
	"github.com/thinhlh/go-web-101/internal/app/domain/product"
	"github.com/thinhlh/go-web-101/internal/core"
)

type IProductRepository interface {
	FindAllProducts() []product.Product
}

type ProductRepository struct {
}

func (ProductRepository) FindAllProducts() []product.Product {
	// var products []product.Product

	type SampleProduct struct {
		ID         int
		Name       string
		Brand      string
		CategoryId int
		Quantity   int
		Price      float64
		Enabled    bool
	}

	var products []SampleProduct
	core.Database.Raw("SELECT * FROM PRODUCT").Scan(&products)

	return make([]product.Product, 0)
}

func NewRepository() IProductRepository {
	return &ProductRepository{}
}
