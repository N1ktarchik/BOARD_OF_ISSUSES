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
	ConnectUserToDesk(ctx context.Context, userID, deskID int, password string) error

	CreateDesk(ctx context.Context, desk *dn.Desk) error
	ChangeDeskName(ctx context.Context, name string, deskId, userID int) error
	ChangeDeskPassword(ctx context.Context, password string, deskId, userID int) error
	ChangeDeskOwner(ctx context.Context, deskId, userID, newOwner int) error
	DeleteDesk(ctx context.Context, deskId, userID int) error
	GetAllDesks(ctx context.Context, userID int) ([]int, error)

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
