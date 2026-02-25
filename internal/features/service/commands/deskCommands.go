package commands

import (
	dn "Board_of_issuses/internal/core/domains"
	"Board_of_issuses/internal/features/service/auth"
	"context"
)

func (s *Service) CreateDesk(ctx context.Context, desk *dn.Desk) error {

	hashPassword, err := auth.Hash(desk.Password)
	if err != nil {
		return err
	}
	desk.Password = hashPassword

	if err := s.repo.CreateDesk(ctx, desk.ToRepoDesk()); err != nil {
		return err
	}

	return nil
}
