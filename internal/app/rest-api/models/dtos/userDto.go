package dtos

import (
	"time"
)

type UserResponse struct {
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GetAllUsersResponse struct {
	Users []*UserResponse `json:"users"`
}
