package api

import "github.com/gin-gonic/gin"

type AuthRouter struct{}

func (r *AuthRouter) Setup(group *gin.RouterGroup) {
	auth := group.Group("/auth")
	auth.POST("/register", r.Register)
	auth.POST("/login-by-email", r.LoginByEmail)
	auth.POST("/login-by-token", r.LoginByToken)
}

func (r *AuthRouter) Register(ctx *gin.Context) {

}

func (r *AuthRouter) LoginByEmail(ctx *gin.Context) {

}

func (r *AuthRouter) LoginByToken(ctx *gin.Context) {

}
