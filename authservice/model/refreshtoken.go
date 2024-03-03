package model

import "time"

type RefreshToken struct {
	Token     string
	User      *User
    ExpiresAt time.Time `db:"expires_at"`
}
