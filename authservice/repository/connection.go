package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresConnection(datasourceURI string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", datasourceURI)
	if err != nil {
		return nil, err
	}
	return db, nil
}
