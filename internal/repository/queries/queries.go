package queries

import (
	"github.com/Reywaltz/web_test/pkg/postgres"
)

type Query struct {
	db *postgres.DB
}

func NewRepository(db *postgres.DB) *Query {
	return &Query{
		db: db,
	}
}
