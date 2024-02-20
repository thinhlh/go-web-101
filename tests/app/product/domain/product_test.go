package domain_test

import (
	"testing"

	product "github.com/thinhlh/go-web-101/internal/app/product/domain"
)

func TestProduct_NewProduct(t *testing.T) {
	type testCase struct {
		test        string
		expectedErr error

		// Data
		name       string
		price      float64
		properties product.ProductProperty
	}

	testCases := []testCase{
		{
			test:  "Create Product Sucessfully",
			name:  "Macbook",
			price: 3000,
			properties: product.ProductProperty{
				Size:  "14 inches",
				Color: "Grey",
			},
			expectedErr: nil,
		},
		{
			test:  "Create Product Unsucessfully when name is empty",
			name:  "",
			price: 3000,
			properties: product.ProductProperty{
				Size:  "14 inches",
				Color: "Grey",
			},
			expectedErr: product.ErrInvalidProductName,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := product.New(tc.name, tc.price, tc.properties)

			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
