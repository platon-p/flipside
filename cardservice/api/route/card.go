package route

import (
	"github.com/gin-gonic/gin"
	"github.com/platon-p/flipside/cardservice/api/controller"
)

type CardRouter struct {
	controller *controller.CardController
}

func NewCardRouter(controller *controller.CardController) *CardRouter {
	return &CardRouter{
		controller: controller,
	}
}

func (r *CardRouter) Setup(group *gin.RouterGroup) {
	cards := group.Group("/cardset/:slug/cards")
	cards.POST("/", r.CreateCards)
	cards.GET("/", r.GetCards)
	cards.PUT("/", r.UpdateCards)
	cards.DELETE("/", r.DeleteCards) // ?positions=...
}

func (r *CardRouter) CreateCards(ctx *gin.Context)
func (r *CardRouter) GetCards(ctx *gin.Context)
func (r *CardRouter) UpdateCards(ctx *gin.Context)
func (r *CardRouter) DeleteCards(ctx *gin.Context)
