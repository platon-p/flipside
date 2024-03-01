package api

import "github.com/gin-gonic/gin"

type Router struct {
	AuthRouter  *AuthRouter
	CheckRouter *CheckRouter
}

func NewRouter(authRouter *AuthRouter, checkRouter *CheckRouter) *Router {
    return &Router{
        AuthRouter: authRouter,
        CheckRouter: checkRouter,
    }
}

func (r *Router) Setup(group *gin.RouterGroup) {
	api := group.Group("/api")

	auth := api.Group("/auth")
	auth.POST("/register", r.AuthRouter.Register)
	auth.POST("/login-by-email", r.AuthRouter.LoginByEmail)
	auth.POST("/login-by-token", r.AuthRouter.LoginByToken)

	check := api.Group("/check")
	check.GET("/email/:email", r.CheckRouter.CheckEmail)
	check.GET("/nickname/:nickname", r.CheckRouter.CheckNickname)
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
