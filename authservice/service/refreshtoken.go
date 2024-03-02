package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/platon-p/flashside/authservice/model"
	"github.com/platon-p/flashside/authservice/repository"
)

var (
	InvalidRefreshToken = errors.New("Invalid refresh token")
	ExpiredRefreshToken = errors.New("Expired refresh token")
)

type RefreshTokenService struct {
	Repository repository.RefreshTokenRepository
	ExpiresIn  time.Duration
}

// Deletes old and returns new
func (s *RefreshTokenService) CreateToken(user *model.User) *model.RefreshToken {
	current := s.Repository.FindByUser(user)
	if current != nil {
		s.Repository.Delete(current.Token)
	}
	tokenStr := s.generateToken()
	if tokenStr == nil {
		return nil
	}
	exp := time.Now().Add(s.ExpiresIn)
	token := s.Repository.Create(user, *tokenStr, exp)
	return token
}

func (s *RefreshTokenService) CheckToken(refreshToken string) (*model.User, error) {
	token := s.Repository.FindByToken(refreshToken)
	if token == nil {
		return nil, InvalidRefreshToken
	}
	if token.ExpiresAt.Before(time.Now()) {
		return nil, ExpiredRefreshToken
	}
	return token.User, nil
}

func (s *RefreshTokenService) generateToken() *string {
	token, err := uuid.NewV7()
	if err != nil {
		return nil
	}
	tokenStr := token.String()
	return &tokenStr
}
