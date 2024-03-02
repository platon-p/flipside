package route

import "github.com/gin-gonic/gin"

type Router struct {
	authRouter  *AuthRouter
	checkRouter *CheckRouter
}

func NewRouter(authRouter *AuthRouter, checkRouter *CheckRouter) *Router {
    return &Router{
        authRouter: authRouter,
        checkRouter: checkRouter,
    }
}

func (r *Router) Setup(group *gin.RouterGroup) {
	api := group.Group("/api")
    r.authRouter.Setup(api)
    r.checkRouter.Setup(api)
}

