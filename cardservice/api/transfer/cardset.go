package transfer

type CardSetRequest struct {
	Title string `json:"title" binding:"required"`
	Slug  string `json:"slug" binding:"required"`
}

type CardSetResponse struct {
	Title   string `json:"title"`
	Slug    string `json:"slug"`
	OwnerId int    `json:"owner_id"`
}

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
