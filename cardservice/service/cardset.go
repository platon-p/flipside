package service

import "github.com/platon-p/flipside/cardservice/model"

type CardSetService struct {
}

func (s *CardSetService) CreateCardSet(ownerId int, cardSet *model.CardSet) error {
	panic("not implemented")
}

func (s *CardSetService) GetCardSet(slug string) (*model.CardSet, error) {
	panic("not implemented")
}

func (s *CardSetService) UpdateCardSet(cardSet *model.CardSet) error {
	panic("not implemented")
}

func (s *CardSetService) DeleteCardSet(cardSetId int) error {
	panic("not implemented")
}
