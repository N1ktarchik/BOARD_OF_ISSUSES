package cache

import (
	"log/slog"

	"github.com/redis/go-redis/v9"
)

type CacheRepository struct {
	db  *redis.Client
	log *slog.Logger
}

func NewCacheRepository(db *redis.Client, log *slog.Logger) *CacheRepository {
	return &CacheRepository{
		db:  db,
		log: log,
	}
}
