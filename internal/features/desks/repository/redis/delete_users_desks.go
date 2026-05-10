package redis

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
)

func (c *DesksCache) DeleteUsersDesks(ctx context.Context, userUUID uuid.UUID) error {
	c.log.Info("deleting all desks from redis", slog.Any("userID", userUUID))

	key := "user_desks_list:" + userUUID.String()

	count := c.db.Del(ctx, key)
	if count.Err() != nil {
		c.log.Error("failed to delete desks from cache", slog.Any("userID", userUUID),
			slog.Any("err", count.Err()))

		return count.Err()
	}

	if count.Val() == 0 {
		c.log.Debug("cache key did not exist", slog.String("key", key))
	} else {
		c.log.Info("successfully invalidated cache", slog.Any("userID", userUUID))
	}

	return nil

}
