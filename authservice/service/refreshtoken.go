package service

import (
	"errors"

	"github.com/platon-p/flashside/authservice/model"
)

var (
	InvalidRefreshToken = errors.New("Invalid refresh token")
	ExpiredToken        = errors.New("Expired refresh token")
)

type RefreshTokenService struct{
}

func (s *RefreshTokenService) CreateToken(user *model.User) string

func (s *RefreshTokenService) CheckToken(token string) *error
