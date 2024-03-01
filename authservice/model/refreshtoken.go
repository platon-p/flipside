package model

import "time"

type RefreshToken struct {
	Token     string
	ExpiresAt time.Time
	User      *User
}
