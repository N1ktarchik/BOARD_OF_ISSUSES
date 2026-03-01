package dto

import (
	dn "Board_of_issuses/internal/core/domains"
	"time"
)

type User struct {
	Id         int       `json:"id"`
	Login      string    `json:"login"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Created_at time.Time `json:"created_at"`
}

func (u *User) ToServiceUser() *dn.User {
	return &dn.User{
		Id:         u.Id,
		Login:      u.Login,
		Password:   u.Password,
		Email:      u.Email,
		Name:       u.Name,
		Created_at: u.Created_at,
	}
}

type UserResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type UpdateNameRequest struct {
	Name string `json:"name"`
}

type UpdateEmailRequest struct {
	Email string `json:"email"`
}

type UpdatePasswordRequest struct {
	Password string `json:"password"`
}

type ConnectUserToDeskRequest struct {
	ID       int    `json:"desk_id"`
	Password string `json:"password"`
}
