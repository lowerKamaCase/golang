package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name string `json:"name" validate:"max=50"`
}

func NewProduct(name string) *Product {
	return &Product{
		Name: name,
	}
}
