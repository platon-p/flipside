package service

import (
	"errors"
	"time"

	"github.com/platon-p/flashside/authservice/model"
	"github.com/platon-p/flashside/authservice/repository"
	"github.com/platon-p/flashside/authservice/utils"
)

var (
	BadCredentialsError = errors.New("Bad credentials")
)

type AuthService struct {
	JwtUtility          *utils.JwtUtility
	PasswordUtility     *utils.PasswordUtility
	UserRepository      repository.UserRepository
	RefreshTokenService RefreshTokenService
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}

func (s *AuthService) Register(name, nickname, email, password string) (*TokenPair, error) {
	passwordHash, err := s.PasswordUtility.GetPasswordHash(password)
	if err != nil {
		return nil, err
	}
	user := model.User{
		Name:     name,
		Nickname: nickname,
		Email:    email,
		Password: *passwordHash,
	}
	res, err := s.UserRepository.Create(&user)
	if err != nil {
		return nil, err
	}
	return s.createTokenPair(res)
}

func (s *AuthService) LoginByEmail(email, password string) (*TokenPair, error) {
	user, err := s.UserRepository.FindByEmail(email)
    if err != nil {
        return nil, err
    }
	if user == nil {
		return nil, BadCredentialsError
	}
	passwordCorrect := s.PasswordUtility.CheckPasswordHash(user.Password, password)
	if !passwordCorrect {
		return nil, BadCredentialsError
	}
	return s.createTokenPair(user)
}

func (s *AuthService) LoginByToken(refreshToken string) (*TokenPair, error) {
	user, err := s.RefreshTokenService.CheckToken(refreshToken)
	if err != nil {
		return nil, err
	}
	return s.createTokenPair(user)
}

func (s *AuthService) createTokenPair(user *model.User) (*TokenPair, error) {
	accessToken, err := s.JwtUtility.CreateAccessToken(*user)
	if err != nil {
		return nil, err
	}
	refreshToken := s.RefreshTokenService.CreateToken(user)

	return &TokenPair{
		AccessToken:  *accessToken,
		RefreshToken: refreshToken.Token,
		ExpiresAt:    refreshToken.ExpiresAt,
	}, nil
}
