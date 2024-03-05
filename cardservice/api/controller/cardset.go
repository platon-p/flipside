package controller

import "github.com/platon-p/flipside/cardservice/api/transfer"

type CardSetController struct {
}

func (r *CardSetController) CreateCardSet(userId int, cardSet *transfer.CardSetRequest) (*transfer.CardSetResponse, error)
func (r *CardSetController) GetCardSet(slug string) (*transfer.CardSetResponse, error)
func (r *CardSetController) UpdateCardSet(userId int, cardSet *transfer.CardSetRequest) (*transfer.CardSetResponse, error)
func (r *CardSetController) DeleteCardSet(userId int, slug string) error
