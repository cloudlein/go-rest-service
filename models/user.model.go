package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" form:"name" validate:"gte=6,lte=32" gorm:"not null"`
	Email     string    `json:"email" form:"email" validate:"required,email" gorm:"not null"`
	Password  string    `json:"password" form:"password" validate:"required,gte=8" gorm:"not null,column:password"`
	Phone     string    `json:"phone" form:"phone" validate:"required,number,min=12" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt" gorm:"index"`
}
