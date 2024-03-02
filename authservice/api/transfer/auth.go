package transfer

import (
	"time"

	"github.com/platon-p/flashside/authservice/service"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type LoginByEmailRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginByTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type TokenPairResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func NewTokenResponse(tokens service.TokenPair) TokenPairResponse {
    return TokenPairResponse{
    	AccessToken:  tokens.AccessToken,
    	RefreshToken: tokens.RefreshToken,
    	ExpiresAt:    tokens.ExpiresAt,
    }
}