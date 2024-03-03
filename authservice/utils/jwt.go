package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/platon-p/flipside/authservice/model"
)

type JwtUtility struct {
	SignAlg   jwt.SigningMethod
	SignKey   interface{}
	ExpiresIn time.Duration
}

func NewJwtUtility(hs256Key string, expiresIn time.Duration) *JwtUtility {
	return &JwtUtility{
		SignAlg:   jwt.SigningMethodHS256,
		SignKey:   hs256Key,
		ExpiresIn: expiresIn,
	}
}

func (u *JwtUtility) CreateAccessToken(user model.User) (*string, error) {
	exp := time.Now().Add(u.ExpiresIn)
	claims := jwt.MapClaims{
		"nickname": user.Nickname,
		"exp":      exp,
	}
	token := jwt.NewWithClaims(u.SignAlg, claims)
	tokenStr, err := token.SignedString(u.SignKey)
	if err != nil {
		return nil, err
	}
	return &tokenStr, nil
}
