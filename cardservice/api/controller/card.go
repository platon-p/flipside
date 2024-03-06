package controller

import (
	"github.com/platon-p/flipside/cardservice/api/transfer"
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
	panic("Not implemented")
}

func (c *CardController) GetCards(slug string) ([]transfer.CardResponse, error) {
	panic("Not implemented")
}

func (c *CardController) UpdateCards(
	userId int,
	slug string,
	request []transfer.CardRequest,
) ([]transfer.CardResponse, error) {
	panic("Not implemented")
}

func (c *CardController) DeleteCards(userId int, slug string, positions []string) error {
	panic("Not implemented")
}
