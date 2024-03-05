package route

import "github.com/gin-gonic/gin"

type CardSetRouter struct {
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
