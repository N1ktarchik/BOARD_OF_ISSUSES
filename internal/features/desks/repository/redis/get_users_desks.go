package redis

import (
	"N1ktarchik/Board_of_issues/internal/core/domain"
	"context"
	"encoding/json"
	"log/slog"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func (c *DesksCache) GetUserDesks(ctx context.Context, userUUID uuid.UUID) ([]domain.Desk, error) {
	c.log.Info("fetching all desks from redis", slog.Any("userID", userUUID))

	key := "user_desks_list:" + userUUID.String()

	data, err := c.db.Get(ctx, key).Result()

	if err != nil {

		if err == redis.Nil {
			c.log.Debug("cache miss for user desks", slog.Any("userID", userUUID))
		} else {
			c.log.Error("redis error during GetUserDesks", slog.Any("err", err))
		}

		return nil, err
	}

	var desks []domain.Desk

	if err := json.Unmarshal([]byte(data), &desks); err != nil {
		c.log.Error("failed to unmarshal desks from cache", slog.Any("userID", userUUID), slog.Any("err", err))
		return nil, err
	}

	c.log.Info("successfully fetched desks from cache", slog.Int("count", len(desks)))
	return desks, err
}
