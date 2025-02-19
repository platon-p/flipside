package service_test

import (
	"github.com/platon-p/flipside/cardservice/model"
	"github.com/platon-p/flipside/cardservice/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockUserRepository struct {
	GetProfileFunc  func(nickname string) (*model.Profile, error)
	GetCardSetsFunc func(userId int) ([]model.CardSet, error)
}

func (m *mockUserRepository) GetProfile(nickname string) (*model.Profile, error) {
	return m.GetProfileFunc(nickname)
}

func (m *mockUserRepository) GetCardSets(userId int) ([]model.CardSet, error) {
	return m.GetCardSetsFunc(userId)
}

func TestUserService_GetProfile(t *testing.T) {
	// not found case
	{
		repo := &mockUserRepository{
			GetProfileFunc: func(nickname string) (*model.Profile, error) {
				return nil, service.ErrProfileNotFound
			},
		}

		s := service.NewUserService(repo)
		_, err := s.GetProfile("nickname")
		assert.ErrorIs(t, err, service.ErrProfileNotFound)
	}

	// positive case
	{
		profile := &model.Profile{Id: 1, Name: "name", Nickname: "nickname"}
		repo := &mockUserRepository{
			GetProfileFunc: func(nickname string) (*model.Profile, error) {
				return profile, nil
			},
		}

		s := service.NewUserService(repo)
		res, err := s.GetProfile(profile.Nickname)
		assert.NoError(t, err)
		assert.Equal(t, profile, res)
	}
}

func TestUserService_GetCardSets(t *testing.T) {
	// profile not found case
	{
		repo := &mockUserRepository{
			GetProfileFunc: func(nickname string) (*model.Profile, error) {
				return nil, service.ErrProfileNotFound
			},
			GetCardSetsFunc: func(userId int) ([]model.CardSet, error) {
				t.Error("unexpected call to GetCardSets")
				return nil, nil
			},
		}
		s := service.NewUserService(repo)
		_, err := s.GetCardSets("nickname")
		assert.ErrorIs(t, err, service.ErrProfileNotFound)
	}

	// no card sets case
	{
		user := &model.Profile{
			Id:       1,
			Name:     "name",
			Nickname: "nickname",
		}
		repo := &mockUserRepository{
			GetProfileFunc: func(nickname string) (*model.Profile, error) {
				return user, nil
			},
			GetCardSetsFunc: func(userId int) ([]model.CardSet, error) {
				return nil, nil
			},
		}
		s := service.NewUserService(repo)
		res, err := s.GetCardSets(user.Nickname)
		assert.NoError(t, err)
		assert.Nil(t, res)
	}

	// card sets found
	{
		user := &model.Profile{
			Id:       1,
			Name:     "name",
			Nickname: "nickname",
		}
		cardSets := []model.CardSet{
			{
				Id:      1,
				Title:   "",
				Slug:    "",
				OwnerId: user.Id,
			},
		}
		repo := &mockUserRepository{
			GetProfileFunc: func(nickname string) (*model.Profile, error) {
				return user, nil
			},
			GetCardSetsFunc: func(userId int) ([]model.CardSet, error) {
				return cardSets, nil
			},
		}
		s := service.NewUserService(repo)
		res, err := s.GetCardSets(user.Nickname)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.ElementsMatch(t, cardSets, res)
	}
}
