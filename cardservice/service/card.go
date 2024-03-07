package service

import (
	"errors"

	"github.com/platon-p/flipside/cardservice/model"
	"github.com/platon-p/flipside/cardservice/repository"
)

var (
	ErrCardNotFound         = errors.New("Card not found")
	ErrNotCardSetOwner      = errors.New("You are not owner of this card set")
	ErrCardNegativePosition = errors.New("Card's position should be positive")
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
	if errors.Is(err, repository.ErrCardSetNotFound) {
		return nil, ErrCardSetNotFound
	}
	if err != nil {
		return nil, err
	}
	if cardSet.OwnerId != userId {
		return nil, ErrNotCardSetOwner
	}
	for i := range cards {
		if cards[i].Position <= 0 {
			return nil, ErrCardNegativePosition
		}
		cards[i].CardSetId = cardSet.Id
	}
	return s.cardRepository.CreateCards(cards)
}

func (s *CardService) UpdateCards(userId int, slug string, cards []model.Card) ([]model.Card, error) {
	cardSet, err := s.cardSetRepository.GetCardSet(slug)
    if errors.Is(err, repository.ErrCardSetNotFound) {
        return nil, ErrCardSetNotFound
    }
	if err != nil {
		return nil, err
	}
	if cardSet.OwnerId != userId {
		return nil, ErrNotCardSetOwner
	}
	for i := range cards {
		cards[i].CardSetId = cardSet.Id
	}
    res, err := s.cardRepository.UpdateCards(cards)
    if errors.Is(err, repository.ErrCardNotFound) {
        return nil, ErrCardNotFound
    }
    if err != nil {
        return nil, err
    }
    return res, nil
}

func (s *CardService) GetCards(slug string) ([]model.Card, error) {
	res, err := s.cardRepository.GetCards(slug)
	if errors.Is(err, repository.ErrCardSetNotFound) {
		return nil, ErrCardSetNotFound
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CardService) DeleteCards(userId int, slug string, positions []int) error {
	cardSet, err := s.cardSetRepository.GetCardSet(slug)
    if errors.Is(err, repository.ErrCardSetNotFound) {
        return ErrCardSetNotFound
    }
	if err != nil {
		return err
	}
	if cardSet.OwnerId != userId {
		return ErrNotCardSetOwner
	}
	err = s.cardRepository.DeleteCards(cardSet.Id, positions)
	if errors.Is(err, repository.ErrCardNotFound) {
		return ErrCardNotFound
	}
	return err
}
