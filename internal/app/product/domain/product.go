package domain

import (
	"errors"

	"github.com/google/uuid"
)

// Root Entity & Aggregate
type Product struct {
	id         uuid.UUID
	name       *string
	price      *float64
	properties *ProductProperty
}

var (
	ErrInvalidProductName = errors.New("invalid Product Name!")
)

func New(name string, price float64, properties ProductProperty) (*Product, error) {

	if name == "" {
		return nil, ErrInvalidProductName
	}
	product := &Product{
		id:         uuid.New(),
		name:       &name,
		price:      &price,
		properties: &properties,
	}

	return product, nil
}
