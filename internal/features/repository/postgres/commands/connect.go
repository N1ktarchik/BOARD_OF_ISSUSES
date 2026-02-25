package commands

import "github.com/jackc/pgx/v5/pgxpool"

type connect struct {
	db *pgxpool.Pool
}
