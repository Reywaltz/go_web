package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type DB struct {
	Session *pgx.Conn
}

func NewDb(cfg Cfg) (*DB, error) {
	conn, err := pgx.Connect(context.Background(), cfg.ConnString)
	if err != nil {
		return nil, fmt.Errorf("%w: Can't connect to db", err)
	}
	return &DB{
		Session: conn}, nil
}
