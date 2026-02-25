package domains

import (
	repo "Board_of_issuses/internal/features/repository"
	"time"
)

type User struct {
	Id         int
	Login      string
	Password   string
	Email      string
	Name       string
	Created_at time.Time
}

func (u *User) ToRepoUser() *repo.User {
	return &repo.User{
		Id:         u.Id,
		Login:      u.Login,
		Password:   u.Password,
		Email:      u.Email,
		Name:       u.Name,
		Created_at: u.Created_at,
	}
}
