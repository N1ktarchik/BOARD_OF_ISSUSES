package postgres

import (
	core_errors "N1ktarchik/Board_of_issues/internal/core/errors"
	"context"
	"log/slog"

	"github.com/google/uuid"
)

func (s *DesksStorage) DeleteDesk(ctx context.Context, userUUID, deskUUID uuid.UUID) error {
	s.log.Info("deleting desk in repository", slog.Any("deskID", deskUUID))

	query := `DELETE FROM desks WHERE id = $1 AND owner_id = $2`

	result, err := s.pool.Exec(ctx, query, deskUUID, userUUID)
	if err != nil {
		s.log.Error("failed to delete desk in repository",
			slog.Any("deskID", deskUUID), slog.Any("userID", userUUID), slog.Any("err", err))
		return core_errors.ServerError()
	}

	if result.RowsAffected() == 0 {
		s.log.Warn("user not owner of desk", slog.Any("userID", userUUID), slog.Any("deskID", deskUUID))

		return core_errors.UserNotOwnerOfDesk(userUUID.String(), deskUUID.String())
	}

	s.log.Info("desk deleted successfully in repository", slog.Any("deskID", deskUUID))

	return nil
}
