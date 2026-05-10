package postgres

import (
	core_errors "N1ktarchik/Board_of_issues/internal/core/errors"
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

func (s *DesksStorage) ConnectUserToDesk(ctx context.Context, userID, deskID uuid.UUID, password string) error {
	s.log.Info("connecting user to desk in repository", slog.Any("deskID", deskID), slog.Any("userID", userID))

	query := `
        INSERT INTO desk_members (user_id, desk_id)
        SELECT $1, id FROM desks 
        WHERE id = $2 AND password = $3`

	result, err := s.pool.Exec(ctx, query, userID, deskID, password)
	if err != nil {

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgForeignKeyViolation {
				s.log.Warn("failed to connect user to desk: desk not found")

				return core_errors.DeskNotFound()
			}
		}

		s.log.Error("failed to connect user to desk: server error",
			slog.Any("userID", userID), slog.Any("deskID", deskID), slog.Any("err", err))

		return core_errors.ServerError()
	}

	if result.RowsAffected() == 0 {
		s.log.Debug("failed to connect user to desk: invalid password or deskID", slog.Any("err", err))
		return core_errors.Forbidden()
	}

	s.log.Info("user connected to desk successfully in repository", slog.Any("deskID", deskID), slog.Any("userID", userID))

	return nil

}
