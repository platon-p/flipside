package route

import "github.com/gin-gonic/gin"

type CardRouter struct {}

func (r *CardRouter) Setup(group *gin.RouterGroup) {
	cards := group.Group("/cardset/:slug/cards")
	cards.POST("/", r.CreateCards)
	cards.GET("/", r.GetCards)
	cards.PUT("/", r.UpdateCards)
	cards.DELETE("/", r.DeleteCards) // ?positions=...
}

func (r *CardSetRouter) CreateCards(ctx *gin.Context)
func (r *CardSetRouter) GetCards(ctx *gin.Context)
func (r *CardSetRouter) UpdateCards(ctx *gin.Context)
func (r *CardSetRouter) DeleteCards(ctx *gin.Context)
