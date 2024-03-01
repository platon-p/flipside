package repository

import "github.com/platon-p/flashside/authservice/model"

type UserRepository interface {
    Create(user *model.User) (*model.User, error)
    FindByEmail(email string) *model.User
    FindByNickname(nickname string) *model.User
}
