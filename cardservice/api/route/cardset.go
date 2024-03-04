package route

import "github.com/gin-gonic/gin"

type CardSetRouter struct {
}

// GET /api/cardset/:slug -> CardSet
// POST /api/cardset (CardSet) -> CardSet
// PUT /api/cardset (CardSet) -> CardSet
// DELETE /api/cardset (CardSet) -> CardSet

// for content
// GET /api/cardset/:slug/cards -> Card
// GET /api/cardset/:slug/card/:id -> Card
// POST /api/cardset/:slug/card (Card) -> Card
// PUT /api/cardset/:slug/card (Card) -> Card
// DELETE /api/cardset/:slug/card (Card) -> Card

// for user
// GET /api/user/:id -> User
// GET /api/user/:id/cardsets -> CardSet

func (c *CardSetRouter) Setup(group *gin.RouterGroup) {
	path := group.Group("/cardset")

	path.GET("/:slug", c.GetCardSet)
	path.POST("/", c.CreateCardSet)
	path.PUT("/", c.UpdateCardSet)
	path.DELETE("/:slug", c.DeleteCardSet)

	path.GET("/:slug/cards", c.GetCardSetCards)
	path.GET("/:slug/cards/:position", c.GetCardSetCard)
	path.POST("/:slug/card", c.CreateCardSetCard)
	path.POST("/:slug/cards", c.CreateCardSetCards)
	path.PUT("/:slug/cards/:position", c.UpdateCardSetCard)
	path.DELETE("/:slug/cards/:position", c.DeleteCardSetCard)
}
