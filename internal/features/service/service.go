package service

import (
	er "Board_of_issuses/internal/core"
	dn "Board_of_issuses/internal/core/domains"
	repo "Board_of_issuses/internal/features/repository"
	"Board_of_issuses/internal/features/service/auth"
	"context"
)

type Service struct {
	repo repo.Repository
	auth auth.Auth
}

func NewService(r repo.Repository, a auth.Auth) *Service {
	return &Service{
		repo: r,
		auth: a,
	}
}

func toRepoUser(user *dn.User) *repo.User {
	return &repo.User{
		Id:         user.Id,
		Login:      user.Login,
		Password:   user.Password,
		Email:      user.Email,
		Name:       user.Name,
		Created_at: user.Created_at,
	}
}

func (s *Service) Registration(ctx context.Context, user *dn.User) (string, error) {
	if user.Login == "" {
		return "", er.NullLogin()
	}

	if user.Password == "" {
		return "", er.NullPassword()
	}

	if user.Name == "" {
		return "", er.NullName()
	}

	register, err := s.repo.CheckUserByEmailAndLogin(ctx, user.Login, user.Email)
	if err == nil {
		return "", err
	}

	if register {
		return "", er.HaveRegister(user.Login)
	}

	hashPassword, err := auth.Hash(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = hashPassword

	if err = s.repo.CreateUser(ctx, toRepoUser(user)); err != nil {
		return "", err
	}

	token, err := s.auth.CreateJwt(user.Id, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) Authorization(ctx context.Context, user *dn.User) (string, error) {
	if user.Login == "" && user.Email == "" {
		return "", er.NullLogin()
	}

	if user.Password == "" {
		return "", er.NullPassword()
	}

	repoUser, err := s.repo.GetUserByLoginOrEmail(ctx, user.Login, user.Email)
	if err != nil {
		return "", err
	}

	if !auth.Compare(user.Password, repoUser.Password) {
		return "", er.InvalidPassword()
	}

	token, err := s.auth.CreateJwt(repoUser.Id, repoUser.Email)
	if err != nil {
		return "", err
	}

	return token, nil

}
