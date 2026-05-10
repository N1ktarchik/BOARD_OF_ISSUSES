package cache

import (
	"N1ktarchik/Board_of_issues/internal/core/domain"
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

func (r *CacheRepository) Get(ctx context.Context, key string) (*domain.IdempotencyRecord, error) {
	data, err := r.db.Get(ctx, key).Result()

	if err == redis.Nil {
		return nil, nil
	}

	if err != nil {
		r.log.Error("redis get error", "err", err, "key", key)
		return nil, err
	}

	result := &domain.IdempotencyRecord{}
	if err := json.Unmarshal([]byte(data), result); err != nil {
		r.log.Error("json unmarshal idempotency record error", "err", err)
		return nil, err
	}

	return result, nil

}

func (r *CacheRepository) Set(ctx context.Context, key string, data *domain.IdempotencyRecord,
	TTL time.Duration) error {

	toCache, err := json.Marshal(data)
	if err != nil {
		r.log.Error("json marshal idempotency record error", "err", err)
		return err
	}

	if err := r.db.Set(ctx, key, toCache, TTL).Err(); err != nil {
		r.log.Error("redis set error", "err", err, "key", key)
		return err
	}

	return nil
}
