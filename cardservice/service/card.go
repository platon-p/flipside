package service

import (
	"errors"

	"github.com/platon-p/flipside/cardservice/model"
	"github.com/platon-p/flipside/cardservice/repository"
)

var (
	ErrNotCardSetOwner = errors.New("You are not owner of this card set")
)

type CardService struct {
	cardSetRepository repository.CardSetRepository
	cardRepository    repository.CardRepository
}

func NewCardService(cardSetRepository repository.CardSetRepository, cardRepository repository.CardRepository) *CardService {
    return &CardService{
    	cardSetRepository: cardSetRepository,
    	cardRepository:    cardRepository,
    }
}

func (s *CardService) CreateCards(userId int, slug string, cards []model.Card) ([]model.Card, error) {
	cardSet, err := s.cardSetRepository.GetCardSet(slug)
	if err != nil {
		return nil, err
	}
	if cardSet == nil {
		return nil, ErrCardSetNotFound
	}
	if cardSet.OwnerId != userId {
		return nil, ErrNotCardSetOwner
	}
	for _, v := range cards {
		v.CardSetId = cardSet.Id
	}
	return s.cardRepository.CreateCards(cards)
}

func (s *CardService) UpdateCards(userId int, slug string, cards []model.Card) ([]model.Card, error) {
	cardSet, err := s.cardSetRepository.GetCardSet(slug)
	if err != nil {
		return nil, err
	}
	if cardSet == nil {
		return nil, ErrCardSetNotFound
	}
	if cardSet.OwnerId != userId {
		return nil, ErrNotCardSetOwner
	}
	for _, v := range cards {
		v.CardSetId = cardSet.Id
	}
	return s.cardRepository.UpdateCards(cards)
}

func (s *CardService) GetCards(slug string) ([]model.Card, error) {
	return s.cardRepository.GetCards(slug)
}

func (s *CardService) DeleteCards(userId int, slug string, positions []int) error {
	cardSet, err := s.cardSetRepository.GetCardSet(slug)
	if err != nil {
		return err
	}
	if cardSet == nil {
		return ErrCardSetNotFound
	}
	if cardSet.OwnerId != userId {
		return ErrNotCardSetOwner
	}
	return s.cardRepository.DeleteCards(slug, positions)
}
