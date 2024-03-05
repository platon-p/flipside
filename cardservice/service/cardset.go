package service

import (
	"github.com/platon-p/flipside/cardservice/model"
	"github.com/platon-p/flipside/cardservice/repository"
)

type CardSetService struct {
	cardSetRepository repository.CardSetRepository
}

func (s *CardSetService) CreateCardSet(cardSet *model.CardSet) (*model.CardSet, error) {
	return s.cardSetRepository.CreateCardSet(cardSet)
}

func (s *CardSetService) GetCardSet(slug string) (*model.CardSet, error) {
	return s.cardSetRepository.GetCardSet(slug)
}

func (s *CardSetService) UpdateCardSet(cardSet *model.CardSet) (*model.CardSet, error) {
	return s.cardSetRepository.UpdateCardSet(cardSet.Id, cardSet)
}

func (s *CardSetService) DeleteCardSet(cardSetId int) error {
	return s.cardSetRepository.DeleteCardSet(cardSetId)
}
