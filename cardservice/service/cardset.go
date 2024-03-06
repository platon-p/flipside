package service

import (
	"errors"

	"github.com/platon-p/flipside/cardservice/model"
	"github.com/platon-p/flipside/cardservice/repository"
)

var (
    ErrCardSetNotFound = errors.New("Card Set not found")
)

type CardSetService struct {
	cardSetRepository repository.CardSetRepository
}

func NewCardSetService(cardSetRepository repository.CardSetRepository) *CardSetService {
    return &CardSetService{
        cardSetRepository: cardSetRepository,
    }
}

func (s *CardSetService) CreateCardSet(cardSet *model.CardSet) (*model.CardSet, error) {
    return s.cardSetRepository.CreateCardSet(cardSet)
}

func (s *CardSetService) GetCardSet(slug string) (*model.CardSet, error) {
    res, err := s.cardSetRepository.GetCardSet(slug)
    if errors.Is(err, repository.ErrCardSetNotFound) {
        return nil, ErrCardSetNotFound
    } 
    if err != nil {
        return nil, err
    }
    return res, nil
}

func (s *CardSetService) UpdateCardSet(cardSet *model.CardSet) (*model.CardSet, error) {
	return s.cardSetRepository.UpdateCardSet(cardSet.Id, cardSet)
}

func (s *CardSetService) DeleteCardSet(ownerId int, cardSetId int) error {
	return s.cardSetRepository.DeleteCardSet(ownerId, cardSetId)
}
