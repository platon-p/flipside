package repository

import (
	"time"

	"github.com/platon-p/flashside/authservice/model"
)

type RefreshTokenRepository interface {
    Create(user *model.User, token string, expiresAt time.Time) *model.RefreshToken
    FindByToken(token string) *model.RefreshToken
    FindByUser(user *model.User) *model.RefreshToken
    Delete(token string) error
}
