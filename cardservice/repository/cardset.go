package repository

import "github.com/platon-p/flipside/cardservice/model"

type CardSetRepository interface {
	CreateCardSet(cardSet *model.CardSet) error
	GetCardSet(slug string) (*model.CardSet, error)
	UpdateCardSet(cardSet *model.CardSet) error
	DeleteCardSet(id int) error
}
