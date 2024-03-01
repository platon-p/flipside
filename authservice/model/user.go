package model

import "time"

type User struct {
	Id        int
	CreatedAt time.Time
	Name      string
	Nickname  string
	Email     string
	Password  string
}
