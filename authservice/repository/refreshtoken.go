package repository

import "github.com/platon-p/flashside/authservice/model"

type RefreshTokenRepository interface {
    Create(user *model.User, token string) *model.RefreshToken
    Find(token string) *model.RefreshToken
    Delete(token string) error
}
