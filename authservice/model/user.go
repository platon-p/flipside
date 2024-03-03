package model

import "time"

type User struct {
	Id        int
    CreatedAt time.Time `db:"created_at"`
	Name      string
	Nickname  string
	Email     string
	Password  string
}
