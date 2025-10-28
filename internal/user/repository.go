package user

import (
	"lowerkamacase/golang/pkg/db"

	"gorm.io/gorm/clause"
)

// type ProductRepositoryDeps struct {
// 	Database *db.Db
// }

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.Database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) Update(user *User) (*User, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) Delete(id uint) error {
	result := repo.Database.DB.Delete(&User{}, "ID = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *UserRepository) GetById(id uint) (*User, error) {
	var user User
	result := repo.Database.DB.First(&user, "ID = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	result := repo.Database.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}