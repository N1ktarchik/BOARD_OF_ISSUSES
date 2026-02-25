package http

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

type Desk struct {
	Id         int
	Name       string
	Password   string
	OwnerId    int
	Created_at time.Time
}

func (d *Desk) ToServiceDeskr() *dn.Desk {
	return &dn.Desk{
		Id:         d.Id,
		Name:       d.Name,
		Password:   d.Password,
		OwnerId:    d.OwnerId,
		Created_at: d.Created_at,
	}

}
