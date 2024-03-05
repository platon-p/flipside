package controller

import (
	"github.com/platon-p/flipside/cardservice/api/transfer"
	"github.com/platon-p/flipside/cardservice/service"
)

type CardSetController struct {
	cardSetService *service.CardSetService
}

func NewCardSetController(cardSetService *service.CardSetService) *CardSetController {
	return &CardSetController{
		cardSetService: cardSetService,
	}
}

func (r *CardSetController) CreateCardSet(userId int, request *transfer.CardSetRequest) (*transfer.CardSetResponse, error) {
	cardSet := request.ToModel(userId)
	newEntity, err := r.cardSetService.CreateCardSet(cardSet)
	if err != nil {
		return nil, err
	}
	return transfer.NewCardSetResponse(newEntity), nil
}

func (r *CardSetController) GetCardSet(slug string) (*transfer.CardSetResponse, error) {
	entity, err := r.cardSetService.GetCardSet(slug)
	if err != nil {
		return nil, err
	}
	return transfer.NewCardSetResponse(entity), nil
}

func (r *CardSetController) UpdateCardSet(userId int, request *transfer.CardSetRequest) (*transfer.CardSetResponse, error) {
	cardSet := request.ToModel(userId)
	newEntity, err := r.cardSetService.UpdateCardSet(cardSet)
	if err != nil {
		return nil, err
	}
	return transfer.NewCardSetResponse(newEntity), nil
}

func (r *CardSetController) DeleteCardSet(userId int, cardSetId int) error {
	return r.cardSetService.DeleteCardSet(userId, cardSetId)
}
