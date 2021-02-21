package queries

import (
	"github.com/Reywaltz/web_test/pkg/postgres"
)

type Query struct {
	db *postgres.DB
}

func New(db *postgres.DB) *Query {
	return &Query{
		db: db,
	}
}
