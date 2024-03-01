package service

import (
	"errors"
	"time"

	"github.com/platon-p/flashside/authservice/model"
	"github.com/platon-p/flashside/authservice/repository"
)

var (
	InvalidRefreshToken = errors.New("Invalid refresh token")
	ExpiredToken        = errors.New("Expired refresh token")
)

type RefreshTokenService struct {
	Repository repository.RefreshTokenRepository
	ExpiresIn  time.Duration
}

// Deletes old and returns new
func (s *RefreshTokenService) CreateToken(user *model.User) string

func (s *RefreshTokenService) CheckToken(token string) (*model.User, error)
