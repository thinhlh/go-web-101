package domain

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Root Entity & Aggregate
type Product struct {
	gorm.Model
	ID         uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name       *string   `gorm:"index"`
	Price      *float64
	Properties *ProductProperty `gorm:"embedded"`
}

var (
	ErrInvalidProductName = errors.New("invalid product name!")
)

func New(name string, price float64, properties ProductProperty) (*Product, error) {

	if name == "" {
		return nil, ErrInvalidProductName
	}
	product := &Product{
		ID:         uuid.New(),
		Name:       &name,
		Price:      &price,
		Properties: &properties,
	}

	return product, nil
}
