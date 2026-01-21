package service

import (
	"errors"

	"github.com/cloudlein/go-rest-service/models"
	"github.com/cloudlein/go-rest-service/repository"
	"github.com/cloudlein/go-rest-service/utils"
)

type UserService interface {
	GetAllUsers(page, limit int) ([]*models.User, int64, error)
	GetUserById(id int64) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func (s *userService) CreateUser(user *models.User) (*models.User, error) {

	if user.Name == "" {
		return nil, errors.New("name is required")
	}

	if user.Email == "" {
		return nil, errors.New("email is required")
	}

	if user.Phone == "" {
		return nil, errors.New("phone is required")
	}

	if user.Password == "" {
		return nil, errors.New("password is required")
	}

	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		Name:     user.Name,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: hashPassword,
	}

	return s.repo.CreateUser(newUser)
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUsers(page, limit int) ([]*models.User, int64, error) {
	return s.repo.FindAll(page, limit)
}

func (s *userService) GetUserById(id int64) (*models.User, error) {
	return s.repo.FindById(id)
}
