package controller

import (
	"github.com/platon-p/flipside/authservice/api/transfer"
	"github.com/platon-p/flipside/authservice/service"
)

type AuthController struct {
	authService  *service.AuthService
	checkService *service.CheckService
}

func NewAuthController(authService *service.AuthService, checkService *service.CheckService) *AuthController {
	return &AuthController{
		authService:  authService,
		checkService: checkService,
	}
}

func (c *AuthController) Register(request transfer.RegisterRequest) (*transfer.TokenPairResponse, error) {
	if err := c.checkService.CheckEmail(request.Email); err != nil {
		return nil, err
	}
	if err := c.checkService.CheckNickname(request.Nickname); err != nil {
		return nil, err
	}
	tokens, err := c.authService.Register(request.Name, request.Nickname, request.Email, request.Password)
	if err != nil {
		return nil, err
	}
	response := transfer.NewTokenPairResponse(*tokens)
	return &response, nil
}

func (c *AuthController) LoginByEmail(request transfer.LoginByEmailRequest) (*transfer.TokenPairResponse, error) {
	tokens, err := c.authService.LoginByEmail(request.Email, request.Password)
	if err != nil {
		return nil, err
	}
	response := transfer.NewTokenPairResponse(*tokens)
	return &response, nil
}

func (c *AuthController) LoginByToken(request transfer.LoginByTokenRequest) (*transfer.TokenPairResponse, error) {
	tokens, err := c.authService.LoginByToken(request.RefreshToken)
	if err != nil {
		return nil, err
	}
	response := transfer.NewTokenPairResponse(*tokens)
	return &response, nil
}
