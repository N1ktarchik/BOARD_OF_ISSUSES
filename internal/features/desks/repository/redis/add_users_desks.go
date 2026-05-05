package redis

import (
	"N1ktarchik/Board_of_issues/internal/core/domain"
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

func (c *DesksCache) SetUserDesks(ctx context.Context, userUUID uuid.UUID, desks []domain.Desk) error {
	c.log.Info("caching desks list", slog.Any("userID", userUUID))

	key := "user_desks_list:" + userUUID.String()

	data, err := json.Marshal(desks)
	if err != nil {
		c.log.Error("failed to marshal desks for cache", slog.Any("err", err))
		return err
	}

	err = c.db.Set(ctx, key, data, 5*time.Hour).Err()
	if err != nil {
		c.log.Error("failed to set desks to redis", slog.Any("err", err))
		return err
	}

	c.log.Info("successfully cached desks", slog.Int("count", len(desks)))

	return nil
}
