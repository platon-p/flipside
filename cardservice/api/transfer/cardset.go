package transfer

import "github.com/platon-p/flipside/cardservice/model"

type CardSetRequest struct {
	Title string `json:"title" binding:"required"`
	Slug  string `json:"slug" binding:"required"`
}

func (r *CardSetRequest) ToModel(ownerId int) *model.CardSet {
	return &model.CardSet{
		Title:   r.Title,
		Slug:    r.Slug,
		OwnerId: ownerId,
	}
}

type CardSetResponse struct {
	Title   string `json:"title"`
	Slug    string `json:"slug"`
	OwnerId int    `json:"owner_id"`
}

func NewCardSetResponse(cardSet *model.CardSet) *CardSetResponse {
	return &CardSetResponse{
		Title:   cardSet.Title,
		Slug:    cardSet.Slug,
		OwnerId: cardSet.OwnerId,
	}
}
