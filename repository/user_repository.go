package repository

import (
	"github.com/cloudlein/go-rest-service/database"
	"github.com/cloudlein/go-rest-service/models"
)

type UserRepository interface {
	FindAll(page, limit int) ([]*models.User, int64, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (u *userRepository) FindAll(page, limit int) ([]*models.User, int64, error) {
	var users []*models.User
	var total int64

	offset := (page - 1) * limit

	if err := database.DB.Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := database.DB.
		Offset(offset).
		Limit(limit).
		Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
