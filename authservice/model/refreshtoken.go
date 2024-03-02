package model

import "time"

type RefreshToken struct {
	Token     string
	User      *User
	ExpiresAt time.Time
}
