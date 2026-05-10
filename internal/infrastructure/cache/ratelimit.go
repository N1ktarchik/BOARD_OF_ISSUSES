package cache

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/redis/go-redis/v9"
)

func (r *CacheRepository) RateLimit(ctx context.Context, userID string,
	limit int, window time.Duration) (bool, error) {

	key := fmt.Sprintf("rl:%s", userID)
	now := time.Now().UnixNano()
	startWindow := now - window.Nanoseconds()

	pipe := r.db.TxPipeline()

	pipe.ZRemRangeByScore(ctx, key, "0", fmt.Sprintf("%d", startWindow))

	count := pipe.ZCard(ctx, key)

	pipe.ZAdd(ctx, key, redis.Z{
		Member: fmt.Sprintf("%d:%d", now, rand.Int63()),
		Score:  float64(now),
	})

	pipe.Expire(ctx, key, window)

	_, err := pipe.Exec(ctx)
	if err != nil {
		r.log.Error("redis rate limit error", "err", err, "key", key)
		return false, err
	}

	return count.Val() < int64(limit), nil
}
