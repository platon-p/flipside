package transfer

import (
	"time"

	"github.com/platon-p/flipside/authservice/service"
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginByEmailRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginByTokenRequest struct {
    RefreshToken string `json:"refresh_token" binding:"required"`
}

type TokenPairResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func NewTokenPairResponse(tokens service.TokenPair) TokenPairResponse {
	return TokenPairResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		ExpiresAt:    tokens.ExpiresAt,
	}
}
