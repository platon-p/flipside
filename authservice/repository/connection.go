package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresConnection(dataSourceURI string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dataSourceURI)
	if err != nil {
		return nil, err
	}
	return db, nil
}
