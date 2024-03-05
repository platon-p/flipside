package controller

import "github.com/platon-p/flipside/cardservice/api/transfer"

type CardController struct {}

func (c *CardController) CreateCards(
    userId int,
    slug string,
    request []transfer.CardRequest,
) ([]transfer.CardResponse, error)

func (c *CardController) GetCards(slug string) ([]transfer.CardResponse, error)

func (c *CardController) UpdateCards(
    slug string,
    request []transfer.CardRequest,
) ([]transfer.CardResponse, error)

func (c *CardController) DeleteCards(slug string, positions []int) error 
