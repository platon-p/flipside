package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewPostgresConnection(dataSourceURI string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceURI)
	if err != nil {
		return nil, err
	}
	return db, nil
}
