package store

import "github.com/jackc/pgx/v5/pgxpool"

type connect struct {
	db *pgxpool.Pool
}

func CreateConnectToDB(db *pgxpool.Pool) *connect {
	return &connect{
		db: db,
	}
}
