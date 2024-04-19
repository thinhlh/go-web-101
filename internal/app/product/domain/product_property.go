package domain

// Value object
type ProductSize string
type ProductColor string

const (
	Red   ProductColor = "Red"
	Green ProductColor = "Green"
	Blue  ProductColor = "Green"
)

const (
	S  ProductSize = "S"
	M  ProductSize = "M"
	L  ProductSize = "L"
	XL ProductSize = "XL"
)

type ProductProperty struct {
	Size  string
	Color string
}
