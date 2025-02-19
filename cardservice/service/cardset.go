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

func (s *CardSetService) UpdateCardSet(oldSlug string, cardSet *model.CardSet) (*model.CardSet, error) {
	oldCardSet, err := s.GetCardSet(oldSlug)
	if err != nil {
		return nil, err
	}
	if oldCardSet.OwnerId != cardSet.OwnerId {
		return nil, ErrNotCardSetOwner
	}
	res, err := s.cardSetRepository.UpdateCardSet(oldSlug, cardSet)
	if errors.Is(err, repository.ErrCardSetNotFound) {
		return nil, ErrCardSetNotFound
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CardSetService) DeleteCardSet(ownerId int, slug string) error {
	err := s.cardSetRepository.DeleteCardSet(ownerId, slug)
	if errors.Is(err, repository.ErrCardSetNotFound) {
		return ErrCardSetNotFound
	}
	return err
}
