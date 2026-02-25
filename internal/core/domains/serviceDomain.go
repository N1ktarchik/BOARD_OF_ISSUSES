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

type Desk struct {
	Id         int
	Name       string
	Password   string
	OwnerId    int
	Created_at time.Time
}

func (d *Desk) ToRepoDesk() *repo.Desk {
	return &repo.Desk{
		Id:         d.Id,
		Name:       d.Name,
		Password:   d.Password,
		OwnerId:    d.OwnerId,
		Created_at: d.Created_at,
	}

}
