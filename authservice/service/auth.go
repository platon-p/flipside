package service

import (
	"errors"
	"time"

	"github.com/platon-p/flipside/authservice/model"
	"github.com/platon-p/flipside/authservice/repository"
	"github.com/platon-p/flipside/authservice/utils"
)

var (
	BadCredentialsError = errors.New("bad credentials")
)

type AuthService struct {
	jwtUtility          *utils.JwtUtility
	passwordUtility     *utils.PasswordUtility
	userRepository      repository.UserRepository
	refreshTokenService *RefreshTokenService
}

func NewAuthService(
	jwtUtility *utils.JwtUtility,
	passwordUtility *utils.PasswordUtility,
	repository repository.UserRepository,
	refreshTokenService *RefreshTokenService,
) *AuthService {
	return &AuthService{
		jwtUtility:          jwtUtility,
		passwordUtility:     passwordUtility,
		userRepository:      repository,
		refreshTokenService: refreshTokenService,
	}
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}

func (s *AuthService) Register(name, nickname, email, password string) (*TokenPair, error) {
	passwordHash, err := s.passwordUtility.GetPasswordHash(password)
	if err != nil {
		return nil, err
	}
	user := model.User{
		Name:     name,
		Nickname: nickname,
		Email:    email,
		Password: *passwordHash,
	}
	res, err := s.userRepository.Create(&user)
	if err != nil {
		return nil, err
	}
	return s.createTokenPair(res)
}

func (s *AuthService) LoginByEmail(email, password string) (*TokenPair, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, BadCredentialsError
	}
	passwordCorrect := s.passwordUtility.CheckPasswordHash(user.Password, password)
	if !passwordCorrect {
		return nil, BadCredentialsError
	}
	return s.createTokenPair(user)
}

func (s *AuthService) LoginByToken(refreshToken string) (*TokenPair, error) {
	user, err := s.refreshTokenService.CheckToken(refreshToken)
	if err != nil {
		return nil, err
	}
	return s.createTokenPair(user)
}

func (s *AuthService) createTokenPair(user *model.User) (*TokenPair, error) {
	accessToken, err := s.jwtUtility.CreateAccessToken(*user)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.refreshTokenService.CreateToken(user)
	if err != nil {
		return nil, err
	}
	return &TokenPair{
		AccessToken:  *accessToken,
		RefreshToken: refreshToken.Token,
		ExpiresAt:    refreshToken.ExpiresAt,
	}, nil
}
