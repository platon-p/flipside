package route

import (
	"github.com/gin-gonic/gin"
	"github.com/platon-p/flipside/cardservice/api/middleware"
)

type Router struct {
	routers []IRouter
}

type IRouter interface {
	Setup(group *gin.RouterGroup)
}

func NewRouter(routers ...IRouter) *Router {
	return &Router{
		routers: routers,
	}
}

func (r *Router) Setup(group *gin.RouterGroup) {
	api := group.Group("/api")
    api.Use(middleware.ErrorHandler())
    for _, v := range r.routers {
        v.Setup(api)
    }
}
