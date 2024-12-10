package transfer

type CardRequest struct {
	Question string `json:"question" binding:"required"`
	Answer   string `json:"answer" binding:"required"`
	Position int    `json:"position" binding:"required"`
}

type CardResponse struct {
	Question  string `json:"question"`
	Answer    string `json:"answer"`
	Position  int    `json:"position"`
	CardSetId int    `json:"card_set_id"`
}

type DeleteCardsRequest struct {
	Query struct {
		Positions []int `query:"positions" binding:"required"`
	}
	Uri struct {
		Slug string `uri:"slug" binding:"required"`
	}
}
