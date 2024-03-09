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
}

func NewUserService(userRepository repository.UserRepository) *UserService {
    return &UserService{
    	userRepository: userRepository,
    }
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
    profile, err := s.GetProfile(nickname)
    if err != nil {
        return nil, err
    }
    res, err := s.userRepository.GetCardSets(profile.Id)
    if errors.Is(err, repository.ErrProfileNotFound) {
        return nil, ErrProfileNotFound
    }
    if err != nil {
        return nil, err
    }
    return res, nil

}
