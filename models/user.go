package models

import (
	user "dot/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
	Products []Product
	Orders   []Order
}

func CoreToModel(data user.CoreUser) User {
	return User{
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
	}
}

func (data *User) ModelToCore() user.CoreUser {
	return user.CoreUser{
		ID:       data.ID,
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
	}
}
