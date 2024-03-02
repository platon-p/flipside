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

func NewRefreshTokenService(repository repository.RefreshTokenRepository, expiresIn time.Duration) *RefreshTokenService {
    return &RefreshTokenService{
    	Repository: repository,
    	ExpiresIn:  expiresIn,
    }
}

// Deletes old and returns new
func (s *RefreshTokenService) CreateToken(user *model.User) (*model.RefreshToken, error) {
	current, err := s.Repository.FindByUser(user.Id)
    if err != nil {
        return nil, err
    }
	if current != nil {
        s.Repository.Delete(current.Token)
	}
	tokenStr, err := s.generateToken()
	if err != nil {
		return nil, err
	}
	exp := time.Now().Add(s.ExpiresIn)
	token, err := s.Repository.Create(user.Id, *tokenStr, exp)
    if err != nil {
        return nil, err
    }
	return token, nil
}

func (s *RefreshTokenService) CheckToken(refreshToken string) (*model.User, error) {
	token, err := s.Repository.FindByToken(refreshToken)
    if err != nil {
        return nil, err
    }
	if token == nil {
		return nil, InvalidRefreshToken
	}
	if token.ExpiresAt.Before(time.Now()) {
		return nil, ExpiredRefreshToken
	}
	return token.User, nil
}

func (s *RefreshTokenService) generateToken() (*string, error) {
	token, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	tokenStr := token.String()
	return &tokenStr, nil
}
