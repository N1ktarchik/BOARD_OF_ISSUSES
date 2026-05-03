package repository

import (
	"N1ktarchik/Board_of_issues/internal/core/domain"
	"context"

	"github.com/google/uuid"
)

func (r *DesksRepository) GetAllUsersDesks(ctx context.Context, userUUID uuid.UUID) ([]domain.Desk, error) {
	cachedDesks, err := r.cache.GetUserDesks(ctx, userUUID)

	if err == nil {
		return cachedDesks, nil
	}

	desks, err := r.storage.GetAllUsersDesks(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	_ = r.cache.SetUserDesks(ctx, userUUID, desks)

	return desks, nil

}
