package controller

import (
	"github.com/platon-p/flipside/cardservice/api/transfer"
	"github.com/platon-p/flipside/cardservice/model"
	"github.com/platon-p/flipside/cardservice/service"
)

type CardController struct {
	cardService *service.CardService
}

func NewCardController(cardService *service.CardService) *CardController {
	return &CardController{
		cardService: cardService,
	}
}

func (c *CardController) CreateCards(
	userId int,
	slug string,
	request []transfer.CardRequest,
) ([]transfer.CardResponse, error) {
	models := make([]model.Card, len(request))
	for i, v := range request {
		models[i] = *cardRequestToModel(&v)
	}
	res, err := c.cardService.CreateCards(userId, slug, models)
	if err != nil {
		return nil, err
	}
	response := make([]transfer.CardResponse, len(res))
	for i, v := range res {
		response[i] = *cardModelToResponse(&v)
	}
	return response, nil
}

func (c *CardController) GetCards(slug string) ([]transfer.CardResponse, error) {
	res, err := c.cardService.GetCards(slug)
	if err != nil {
		return nil, err
	}
	response := make([]transfer.CardResponse, len(res))
	for i, v := range res {
		response[i] = *cardModelToResponse(&v)
	}
	return response, nil
}

func (c *CardController) UpdateCards(
	userId int,
	slug string,
	request []transfer.CardRequest,
) ([]transfer.CardResponse, error) {
	models := make([]model.Card, len(request))
	for i, v := range request {
		models[i] = *cardRequestToModel(&v)
	}
	res, err := c.cardService.UpdateCards(userId, slug, models)
	if err != nil {
		return nil, err
	}
	response := make([]transfer.CardResponse, len(res))
	for i, v := range res {
		response[i] = *cardModelToResponse(&v)
	}
	return response, nil
}

func (c *CardController) DeleteCards(userId int, slug string, positions []string) error {
	return c.DeleteCards(userId, slug, positions)
}

func cardRequestToModel(request *transfer.CardRequest) *model.Card {
	return &model.Card{
		Question: request.Question,
		Answer:   request.Answer,
		Position: request.Position,
	}
}

func cardModelToResponse(card *model.Card) *transfer.CardResponse {
	return &transfer.CardResponse{
		Question:  card.Question,
		Answer:    card.Answer,
		Position:  card.Position,
		CardSetId: card.CardSetId,
	}
}
