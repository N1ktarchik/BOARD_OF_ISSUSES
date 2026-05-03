package postgres

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	pgForeignKeyViolation = "23503"
)

type DesksStorage struct {
	pool *pgxpool.Pool
	log  *slog.Logger
}

func NewDesksStorage(pool *pgxpool.Pool, log *slog.Logger) *DesksStorage {
	return &DesksStorage{
		pool: pool,
		log:  log,
	}
}
