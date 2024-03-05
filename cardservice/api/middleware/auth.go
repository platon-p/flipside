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

func (m *AuthMiddleware) ValidateToken(token string) (*int, error) {
	var claims jwt.MapClaims
	jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return m.SignKey, nil
	})
	userId, ok := claims["userId"].(int)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}
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
