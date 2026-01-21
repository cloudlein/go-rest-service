package repository

import (
	"github.com/cloudlein/go-rest-service/database"
	"github.com/cloudlein/go-rest-service/models"
)

type UserRepository interface {
	FindAll(page, limit int) ([]*models.User, int64, error)
	FindById(id int64) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
}

type userRepository struct{}

func (u *userRepository) CreateUser(user *models.User) (*models.User, error) {
	if err := database.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

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

func (u *userRepository) FindById(id int64) (*models.User, error) {
	var user models.User

	if err := database.DB.Model(&models.User{}).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
