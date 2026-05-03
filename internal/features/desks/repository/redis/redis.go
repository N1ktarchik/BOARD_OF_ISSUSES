package redis

import (
	"log/slog"

	redis "github.com/redis/go-redis/v9"
)

type DesksCache struct {
	db  *redis.Client
	log *slog.Logger
}

func NewDesksCache(db *redis.Client, log *slog.Logger) *DesksCache {
	return &DesksCache{
		db:  db,
		log: log,
	}
}
