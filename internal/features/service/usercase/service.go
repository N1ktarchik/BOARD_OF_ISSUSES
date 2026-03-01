package usercase

import (
	repo "Board_of_issuses/internal/features/repository"
	auth "Board_of_issuses/internal/features/service/authJWT"
)

type Service struct {
	repo repo.Repository
	auth *auth.AuthManager
}

func NewService(r repo.Repository, a *auth.AuthManager) *Service {
	return &Service{
		repo: r,
		auth: a,
	}
}
