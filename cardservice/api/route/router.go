package route

import "github.com/gin-gonic/gin"

type Router struct {
	cardSetRouter *CardSetRouter
	cardRouter    *CardRouter
}

func NewRouter(cardSetRouter *CardSetRouter, cardRouter *CardRouter) *Router {
	return &Router{
		cardSetRouter: cardSetRouter,
		cardRouter:    cardRouter,
	}
}

func (r *Router) Setup(group *gin.RouterGroup) {
	api := group.Group("/api")
	r.cardSetRouter.Setup(api)
	r.cardRouter.Setup(api)
}
