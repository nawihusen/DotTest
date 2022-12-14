package controller

import user "dot/features/users"

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func CoreToResUser(core user.CoreUser) UserResponse {
	return UserResponse{
		Username: core.Username,
		Email:    core.Email,
	}
}
