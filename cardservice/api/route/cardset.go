package route

import (
	"github.com/gin-gonic/gin"
	"github.com/platon-p/flipside/cardservice/api/controller"
	"github.com/platon-p/flipside/cardservice/api/transfer"
)

type CardSetRouter struct {
    controller *controller.CardSetController
}

func (r *CardSetRouter) Setup(group *gin.RouterGroup) {
	cardset := group.Group("/cardset")

	cardset.POST("/", r.CreateCardSet)
	cardset.GET("/:slug", r.GetCardSet)
	cardset.PUT("/", r.UpdateCardSet)
	cardset.DELETE("/:slug", r.DeleteCardSet)
}

func (r *CardSetRouter) CreateCardSet(ctx *gin.Context)
func (r *CardSetRouter) GetCardSet(ctx *gin.Context)
func (r *CardSetRouter) UpdateCardSet(ctx *gin.Context)
func (r *CardSetRouter) DeleteCardSet(ctx *gin.Context)
