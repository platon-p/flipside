package data

import "github.com/platon-p/flashside/authservice/model"

type UserRepository interface {
    Create(user *model.User) *model.User
    FindByEmail(email string) *model.User
    FindByNickname(nickname string) *model.User
}
