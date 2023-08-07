package database

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Database struct {
	db *pgxpool.Pool
}

func NewConnection(connString string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
