package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/platon-p/flipside/cardservice/api/helper"
)

type AuthMiddleware struct {
	SignKey interface{}
}

func NewAuthMiddleware(SignKey interface{}) *AuthMiddleware {
    return &AuthMiddleware{
    	SignKey: SignKey,
    }
}

func (m *AuthMiddleware) ValidateToken(token string) (*int, error) {
	var claims jwt.MapClaims
    _, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return m.SignKey, nil
	}, jwt.WithExpirationRequired())
    if err != nil {
        return nil, err
    }
	userIdFloat, ok := claims["id"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}
    userId := int(userIdFloat)
	return &userId, nil
}

func (m *AuthMiddleware) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		token, found := strings.CutPrefix(authHeader, "Bearer ")
		if !found {
			helper.ErrorMessage(ctx, http.StatusUnauthorized, helper.Unauthorized)
			ctx.Abort()
			return
		}
		userId, err := m.ValidateToken(token)
		if err != nil {
			fmt.Println("Auth middleware:", err)
			helper.ErrorMessage(ctx, http.StatusUnauthorized, helper.Unauthorized)
			ctx.Abort()
			return
		}
		ctx.Set("userId", userId)
		ctx.Next()
	}
}
