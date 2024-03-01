package api

import "github.com/gin-gonic/gin"

func AddRoutes(group *gin.RouterGroup, authRouter AuthRouter, checkRouter CheckRouter) {
	api := group.Group("/api")

	auth := api.Group("/auth")
	auth.POST("/register", authRouter.Register)
	auth.POST("/login-by-email", authRouter.LoginByEmail)
	auth.POST("/login-by-token", authRouter.LoginByToken)

	check := api.Group("/check")
	check.GET("/email/:email", checkRouter.CheckEmail)
	check.GET("/nickname/:nickname", checkRouter.CheckNickname)
}

type AuthRouter struct{}

func (r *AuthRouter) Register(ctx *gin.Context) {

}

func (r *AuthRouter) LoginByEmail(ctx *gin.Context) {

}

func (r *AuthRouter) LoginByToken(ctx *gin.Context) {

}

type CheckRouter struct{}

func (r *CheckRouter) CheckEmail(ctx *gin.Context) {

}

func (r *CheckRouter) CheckNickname(ctx *gin.Context) {

}
