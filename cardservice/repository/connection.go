package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewConnection(dataSource string) (*sqlx.DB, error) {
    db, err := sqlx.Connect("postgres", dataSource)
    return db, err
}
