package service

import (
	"errors"

	"github.com/platon-p/flashside/authservice/model"
	"github.com/platon-p/flashside/authservice/repository"
	"github.com/platon-p/flashside/authservice/utils"
)

var (
    BadCredentials = errors.New("Bad credentials")
)

type AuthService struct {
	JwtUtility          *utils.JwtUtility
	PasswordUtility     *utils.PasswordUtility
	UserRepository      repository.UserRepository
	RefreshTokenService RefreshTokenService
}

type TokenPair struct {
	AccessToken  string
	TokenType    string
	RefreshToken string
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
	accessToken, err := s.JwtUtility.CreateAccessToken(*res)
	if err != nil {
		return nil, err
	}
	refreshToken := s.RefreshTokenService.CreateToken(res)

    return &TokenPair{
    	AccessToken:  *accessToken,
    	TokenType:    "bearer",
    	RefreshToken: refreshToken,
    }, nil
}

func (s *AuthService) LoginByEmail(email, password string) (*TokenPair, error) {
    user := s.UserRepository.FindByEmail(email)
    if user == nil {
        return nil, BadCredentials
    }

}

func (s *AuthService) LoginByToken(refreshToken string) (*TokenPair, error)
