package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateProduct struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewProduct(name string, description string) *Product {
	return &Product{
		Name:        name,
		Description: description,
	}
}
