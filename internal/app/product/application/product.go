package application

import (
	"github.com/thinhlh/go-web-101/internal/app/product/domain"
	"github.com/thinhlh/go-web-101/internal/app/product/infrastructure"
)

type ProductService interface {
	GetAllProducts() []domain.Product
	GetProductById() domain.Product
}

type ProductServiceImpl struct {
	repository infrastructure.ProductRepository
}

func NewProductService(repository infrastructure.ProductRepository) ProductServiceImpl {
	return ProductServiceImpl{repository: repository}
}

func (s ProductServiceImpl) GetAllProducts() []domain.Product {
	return s.repository.GetAllProducts()
}

func (s ProductServiceImpl) GetProductById() domain.Product {
	return s.repository.GetProductById()
}
