package web

import "login-management/model/domain"

type UserSessionResponse struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Token     string `json:"token"`
}

func NewUserSessionResponse(user domain.User, token string) UserSessionResponse {
	return UserSessionResponse{
		Id:        user.Id,
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Token:     token,
	}
}
