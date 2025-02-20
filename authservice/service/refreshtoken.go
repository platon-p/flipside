package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/platon-p/flipside/authservice/model"
	"github.com/platon-p/flipside/authservice/repository"
)

var (
	InvalidRefreshToken = errors.New("invalid refresh token")
	ExpiredRefreshToken = errors.New("expired refresh token")
)

type RefreshTokenService struct {
	repository repository.RefreshTokenRepository
	expiresIn  time.Duration
}

func NewRefreshTokenService(repository repository.RefreshTokenRepository, expiresIn time.Duration) *RefreshTokenService {
	return &RefreshTokenService{
		repository: repository,
		expiresIn:  expiresIn,
	}
}

// CreateToken creates a new refresh token for the user
// If there is an existing token, it will be deleted
// Assumes that the user exists and not nil
func (s *RefreshTokenService) CreateToken(user *model.User) (*model.RefreshToken, error) {
	oldToken, err := s.repository.FindByUser(user.Id)
	if err != nil {
		return nil, err
	}
	if oldToken != nil {
		s.repository.Delete(oldToken.Token)
	}
	tokenStr, err := s.generateToken()
	if err != nil {
		return nil, err
	}
	exp := time.Now().Add(s.expiresIn)
	token, err := s.repository.Create(user.Id, *tokenStr, exp)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// CheckToken checks if the refresh token is valid
// if the token is invalid, it returns InvalidRefreshToken
// if the token is expired, it returns ExpiredRefreshToken
// if the token is valid, it returns the user associated with the token
func (s *RefreshTokenService) CheckToken(refreshToken string) (*model.User, error) {
	token, err := s.repository.FindByToken(refreshToken)
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
