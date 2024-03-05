package route

import "github.com/gin-gonic/gin"

type CardSetRouter struct {
}

// for user
// GET /api/user/:id -> User
// GET /api/user/:id/cardsets -> CardSet

func (c *CardSetRouter) Setup(group *gin.RouterGroup) {
	cardset := group.Group("/cardset")

	cardset.POST("/", c.CreateCardSet)
	cardset.GET("/:slug", c.GetCardSet)
	cardset.PUT("/", c.UpdateCardSet)
	cardset.DELETE("/:slug", c.DeleteCardSet)

    cards := cardset.Group("/:slug/cards")
    cards.POST("/", c.CreateCards)
    cards.GET("/", c.GetCards)
    cards.PUT("/", c.UpdateCards)
    cards.DELETE("/", c.DeleteCards) // ?positions=...
}
