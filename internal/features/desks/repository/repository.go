package repository

import (
	"N1ktarchik/Board_of_issues/internal/core/domain"
	"N1ktarchik/Board_of_issues/internal/features/desks/repository/postgres"
	"N1ktarchik/Board_of_issues/internal/features/desks/repository/redis"
	"context"

	"github.com/google/uuid"
)

type DesksRepository struct {
	storage *postgres.DesksStorage
	cache   *redis.DesksCache
}

func NewDesksRepository(storage *postgres.DesksStorage, cache *redis.DesksCache) *DesksRepository {
	return &DesksRepository{
		storage: storage,
		cache:   cache,
	}
}

func (r *DesksRepository) CreateDesk(ctx context.Context, desk *domain.Desk) (*domain.Desk, error) {
	result, err := r.storage.CreateDesk(ctx, desk)
	if err != nil {
		return result, err
	}

	_ = r.cache.DeleteUsersDesks(ctx, desk.OwnerId)
	return result, nil
}

func (r *DesksRepository) ChangeDesksData(ctx context.Context, deskUpdate *domain.Desk,
	requesterID uuid.UUID) (*domain.Desk, error) {

	result, err := r.storage.ChangeDesksData(ctx, deskUpdate, requesterID)
	if err != nil {
		return result, err
	}

	_ = r.cache.DeleteUsersDesks(ctx, deskUpdate.OwnerId)
	return result, nil
}

func (r *DesksRepository) DeleteDesk(ctx context.Context, userUUID, deskUUID uuid.UUID) error {
	if err := r.storage.DeleteDesk(ctx, userUUID, deskUUID); err != nil {
		return err
	}

	_ = r.cache.DeleteUsersDesks(ctx, userUUID)
	return nil
}

func (r *DesksRepository) ConnectUserToDesk(ctx context.Context, userID, deskID uuid.UUID,
	password string) error {
	if err := r.storage.ConnectUserToDesk(ctx, userID, deskID, password); err != nil {
		return err
	}

	_ = r.cache.DeleteUsersDesks(ctx, userID)
	return nil
}
