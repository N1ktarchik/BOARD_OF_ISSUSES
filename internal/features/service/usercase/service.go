package usercase

import (
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

func (s *Service) CreateJWT(ctx context.Context, jwtToken string) (int, error) {
	claims, err := s.auth.Validate(jwtToken)
	if err != nil {
		return 0, err
	}

	return claims.UserId, nil

}
