package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
}

func (db *DB) Pool() *pgxpool.Pool {
	return db.pool
}

func NewDb(cfg Cfg) (*DB, error) {
	conn, err := pgxpool.Connect(context.Background(), cfg.ConnString)
	if err != nil {
		return nil, fmt.Errorf("%w: Can't connect to db", err)
	}
	return &DB{
		pool: conn}, nil
}
