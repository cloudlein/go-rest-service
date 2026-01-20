package service

import (
	"github.com/cloudlein/go-rest-service/models"
	"github.com/cloudlein/go-rest-service/repository"
)

type UserService interface {
	GetAllUsers(page, limit int) ([]*models.User, int64, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUsers(page, limit int) ([]*models.User, int64, error) {
	return s.repo.FindAll(page, limit)
}
