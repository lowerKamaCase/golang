package product

import (
	"lowerkamacase/golang/pkg/db"

	"gorm.io/gorm/clause"
)

type ProductRepositoryDeps struct {
	Database *db.Db
}

type ProductRepository struct {
	Database *db.Db
}

func NewProductRepository(database *db.Db) *ProductRepository {
	return &ProductRepository{
		Database: database,
	}
}

func (repo *ProductRepository) Create(product *Product) (*Product, error) {
	result := repo.Database.DB.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *ProductRepository) Update(product *Product) (*Product, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *ProductRepository) Delete(id uint) error {
	result := repo.Database.DB.Delete(&Product{}, "ID = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *ProductRepository) GetById(id uint) (*Product, error) {
	var product Product
	result := repo.Database.DB.First(&product, "ID = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}