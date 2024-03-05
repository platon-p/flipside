package controller

import "github.com/platon-p/flipside/cardservice/api/transfer"

type CardController struct{}

func NewCardController() *CardController {
	return &CardController{}
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
	slug string,
	request []transfer.CardRequest,
) ([]transfer.CardResponse, error) {
	panic("Not implemented")
}

func (c *CardController) DeleteCards(slug string, positions []int) error {
	panic("Not implemented")
}
