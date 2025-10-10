package product

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Images pq.StringArray `json:"images" gorm:"type:text[]"`
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
