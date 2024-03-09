package service

import (
	"errors"

	"github.com/platon-p/flipside/cardservice/model"
	"github.com/platon-p/flipside/cardservice/repository"
)

var (
    ErrProfileNotFound = errors.New("Profile not found")
)
type UserService struct {
	userRepository    repository.UserRepository
	cardSetRepository repository.CardSetRepository
}

func (s *UserService) GetProfile(nickname string) (*model.Profile, error) {
    res, err := s.userRepository.GetProfile(nickname)
    if errors.Is(err, repository.ErrProfileNotFound) {
        return nil, ErrProfileNotFound
    }
    if err != nil {
        return nil, err
    }
    return res, nil
}

func (s *UserService) GetCardSets(nickname string) ([]model.CardSet, error) {

}
