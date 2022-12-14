package controller

import (
	user "dot/features/users"
)

type UserRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UpdateRequest struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (client *UserRequest) ReqToCore() user.CoreUser {
	return user.CoreUser{
		Username: client.Username,
		Email:    client.Email,
		Password: client.Password,
	}
}

func (client *LoginRequest) ReqToCoreLogin() user.CoreUser {
	return user.CoreUser{
		Email:    client.Email,
		Password: client.Password,
	}
}

func (client *UpdateRequest) ReqToCoreUpdate() user.CoreUser {
	return user.CoreUser{
		Username: client.Username,
		Email:    client.Email,
		Password: client.Password,
	}
}
