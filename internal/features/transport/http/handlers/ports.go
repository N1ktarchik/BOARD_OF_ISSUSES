package handlers

import (
	dn "Board_of_issuses/internal/core/domains"
	"context"
)

type Service interface {
	Registration(ctx context.Context, user *dn.User) (string, error)
	Authorization(ctx context.Context, user *dn.User) (string, error)

	ChangeUserName(ctx context.Context, name string, userID int) error
	ChangeUserEmail(ctx context.Context, email string, userID int) error
	ChangeUserPassword(ctx context.Context, password string, userID int) error

	CreateDesk(ctx context.Context, desk *dn.Desk) error

	CreateJWT(ctx context.Context, jwtToken string) (int, error)
}

type UserHandler struct {
	serv Service
}

func NewUserHandler(src Service) *UserHandler {
	return &UserHandler{
		serv: src,
	}
}
