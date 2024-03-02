package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewPostgresConnection(datasourceURI string) (*sql.DB, error) {
	db, err := sql.Open("postgres", datasourceURI)
	if err != nil {
		return nil, err
	}
	return db, nil
}
