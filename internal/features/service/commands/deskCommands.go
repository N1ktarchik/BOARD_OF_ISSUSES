package commands

import (
	er "Board_of_issuses/internal/core"
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

func (s *Service) ChangeDeskName(ctx context.Context, name string, deskId, userID int) error {
	ownerID, err := s.repo.CheckDeskOwner(ctx, deskId)
	if err != nil {
		return err
	}

	if ownerID != userID {
		return er.UserNotOwnerOfDesk(userID, deskId)
	}

	if err := s.repo.UpdateDeskName(ctx, deskId, name); err != nil {
		return err
	}

	return nil
}

func (s *Service) ChangeDeskPassword(ctx context.Context, password string, deskId, userID int) error {
	ownerID, err := s.repo.CheckDeskOwner(ctx, deskId)
	if err != nil {
		return err
	}

	if ownerID != userID {
		return er.UserNotOwnerOfDesk(userID, deskId)
	}

	hashPass, err := auth.Hash(password)
	if err != nil {
		return err
	}

	if err := s.repo.UpdateDesksPassword(ctx, deskId, hashPass); err != nil {
		return err
	}

	return nil
}

func (s *Service) ChangeDeskOwner(ctx context.Context, deskId, userID, newOwner int) error {
	ownerID, err := s.repo.CheckDeskOwner(ctx, deskId)
	if err != nil {
		return err
	}

	if ownerID != userID {
		return er.UserNotOwnerOfDesk(userID, deskId)
	}

	if err := s.repo.UpdateDeskOwner(ctx, newOwner, deskId); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteDesk(ctx context.Context, deskId, userID int) error {
	ownerID, err := s.repo.CheckDeskOwner(ctx, deskId)
	if err != nil {
		return err
	}

	if ownerID != userID {
		return er.UserNotOwnerOfDesk(userID, deskId)
	}

	if err := s.repo.DeleteDesk(ctx, deskId); err != nil {
		return err
	}

	return nil
}
