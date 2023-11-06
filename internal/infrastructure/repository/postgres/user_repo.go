package postgres

import (
	"go-hexagonal/internal/domain"
	"go-hexagonal/internal/ports/output"

	"gorm.io/gorm"
)

var _ output.UserRepositoryPort = &UserRepository{}

type UserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) domain.UserRepository {
	return &UserRepository{
		db,
	}
}

func (r *UserRepository) Save(user domain.User) error {
	return r.db.Create(&user).Error
}

func (r *UserRepository) FindByID(id string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(user domain.User) error {
	return r.db.Save(&user).Error
}

func (r *UserRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&domain.User{}).Error
}
