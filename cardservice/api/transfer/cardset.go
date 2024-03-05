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
