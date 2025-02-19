package service_test

import (
	"github.com/platon-p/flipside/cardservice/model"
	"github.com/platon-p/flipside/cardservice/repository"
	"github.com/platon-p/flipside/cardservice/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockCardSetRepository struct {
	CreateCardSetFunc func(cardSet *model.CardSet) (*model.CardSet, error)
	GetCardSetFunc    func(slug string) (*model.CardSet, error)
	UpdateCardSetFunc func(oldSlug string, cardSet *model.CardSet) (*model.CardSet, error)
	DeleteCardSetFunc func(ownerId int, slug string) error
}

func (m *mockCardSetRepository) CreateCardSet(cardSet *model.CardSet) (*model.CardSet, error) {
	return m.CreateCardSetFunc(cardSet)
}

func (m *mockCardSetRepository) GetCardSet(slug string) (*model.CardSet, error) {
	return m.GetCardSetFunc(slug)
}

func (m *mockCardSetRepository) UpdateCardSet(oldSlug string, cardSet *model.CardSet) (*model.CardSet, error) {
	return m.UpdateCardSetFunc(oldSlug, cardSet)
}

func (m *mockCardSetRepository) DeleteCardSet(ownerId int, slug string) error {
	return m.DeleteCardSetFunc(ownerId, slug)
}

func TestCardSetService_GetCardSet(t *testing.T) {
	// check not found
	{
		repo := &mockCardSetRepository{
			GetCardSetFunc: func(slug string) (*model.CardSet, error) {
				return nil, repository.ErrCardSetNotFound
			},
		}
		s := service.NewCardSetService(repo)
		_, err := s.GetCardSet("slug")
		assert.ErrorIs(t, err, service.ErrCardSetNotFound)
	}
	// found
	{
		cardSet := &model.CardSet{Slug: "slug"}
		repo := &mockCardSetRepository{
			GetCardSetFunc: func(slug string) (*model.CardSet, error) {
				return cardSet, nil
			},
		}
		s := service.NewCardSetService(repo)
		res, err := s.GetCardSet("slug")
		assert.NoError(t, err)
		assert.Equal(t, cardSet, res)
	}
}

func TestCardSetService_UpdateCardSet(t *testing.T) {
	// check not found
	{
		repo := &mockCardSetRepository{
			GetCardSetFunc: func(slug string) (*model.CardSet, error) {
				return nil, repository.ErrCardSetNotFound
			},
		}
		s := service.NewCardSetService(repo)
		_, err := s.UpdateCardSet("slug", &model.CardSet{})
		assert.ErrorIs(t, err, service.ErrCardSetNotFound)
	}
	// check not owner
	{
		repo := &mockCardSetRepository{
			GetCardSetFunc: func(slug string) (*model.CardSet, error) {
				return &model.CardSet{OwnerId: 1}, nil
			},
		}
		s := service.NewCardSetService(repo)
		_, err := s.UpdateCardSet("slug", &model.CardSet{OwnerId: 2})
		assert.ErrorIs(t, err, service.ErrNotCardSetOwner)
	}
	// found
	{
		cardSet := &model.CardSet{Slug: "slug", OwnerId: 1}
		repo := &mockCardSetRepository{
			GetCardSetFunc: func(slug string) (*model.CardSet, error) {
				if slug != cardSet.Slug {
					return nil, repository.ErrCardSetNotFound
				}
				return cardSet, nil
			},
			UpdateCardSetFunc: func(oldSlug string, arg *model.CardSet) (*model.CardSet, error) {
				if oldSlug != cardSet.Slug {
					return nil, repository.ErrCardSetNotFound
				}
				*cardSet = *arg
				return cardSet, nil
			},
		}
		s := service.NewCardSetService(repo)
		newCardSet := &model.CardSet{Slug: "slug2", OwnerId: 1}
		res, err := s.UpdateCardSet("slug", newCardSet)
		assert.NoError(t, err)
		assert.Equal(t, newCardSet, res)
		_, err = s.GetCardSet("slug")
		assert.ErrorIs(t, err, service.ErrCardSetNotFound)
	}
}

func TestCardSetService_DeleteCardSet(t *testing.T) {
	// check not found
	{
		repo := &mockCardSetRepository{
			DeleteCardSetFunc: func(ownerId int, slug string) error {
				return repository.ErrCardSetNotFound
			},
		}
		s := service.NewCardSetService(repo)
		err := s.DeleteCardSet(1, "slug")
		assert.ErrorIs(t, err, service.ErrCardSetNotFound)
	}
	// found
	{
		repo := &mockCardSetRepository{
			DeleteCardSetFunc: func(ownerId int, slug string) error {
				return nil
			},
		}
		s := service.NewCardSetService(repo)
		err := s.DeleteCardSet(1, "slug")
		assert.NoError(t, err)
	}
}
