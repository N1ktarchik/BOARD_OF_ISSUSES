package usercase

import (
	er "Board_of_issuses/internal/core"
	"Board_of_issuses/internal/core/auth"
	dn "Board_of_issuses/internal/core/domains"
	"context"
)

func (s *Service) accessVerificationToDeskForUser(ctx context.Context, userID, deskID int) error {
	access, err := s.repo.CheckUserDesk(ctx, userID, deskID)
	if err != nil {
		return err
	}

	if !access {
		return er.UserHaveNotAccesToDesk(userID, deskID)
	}

	return nil
}

func (s *Service) accessVerificationToDeskForOwner(ctx context.Context, userID, deskID int) error {
	ownerID, err := s.repo.CheckDeskOwner(ctx, deskID)
	if err != nil {
		return err
	}

	if ownerID != userID {
		return er.UserNotOwnerOfDesk(userID, deskID)
	}

	return nil
}

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
	if err := s.accessVerificationToDeskForOwner(ctx, userID, deskId); err != nil {
		return err
	}

	if err := s.repo.UpdateDeskName(ctx, deskId, name); err != nil {
		return err
	}

	return nil
}

func (s *Service) ChangeDeskPassword(ctx context.Context, password string, deskId, userID int) error {
	if err := s.accessVerificationToDeskForOwner(ctx, userID, deskId); err != nil {
		return err
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
	if err := s.accessVerificationToDeskForOwner(ctx, userID, deskId); err != nil {
		return err
	}

	if err := s.repo.UpdateDeskOwner(ctx, newOwner, deskId); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteDesk(ctx context.Context, deskId, userID int) error {
	if err := s.accessVerificationToDeskForOwner(ctx, userID, deskId); err != nil {
		return err
	}

	if err := s.repo.DeleteDesk(ctx, deskId); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetAllDesks(ctx context.Context, userID int) ([]int, error) {
	return s.repo.GetUserDesks(ctx, userID)
}
