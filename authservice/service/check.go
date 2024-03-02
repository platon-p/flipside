package service

import (
	"errors"

	"github.com/platon-p/flashside/authservice/repository"
)

var (
	EmailIncorrectFormatError = errors.New("Incorrect format of email")
	EmailExistsError          = errors.New("Email already exists")

	NicknameIncorrectFormatError = errors.New("Incorect format of nickname")
	NicknameExistsError          = errors.New("Nickname already exists")
)

type CheckService struct {
	userRepository repository.UserRepository
}

func (s *CheckService) CheckEmail(email string) error {
    found, err := s.userRepository.FindByEmail(email)
    if err != nil {
        return err
    }
    if found == nil {
        return nil
    }
    return EmailExistsError
}

func (s *CheckService) CheckNickname(nickname string) error {
    found, err := s.userRepository.FindByNickname(nickname)
    if err != nil {
        return nil
    }
    if found == nil {
        return nil
    }
    return NicknameExistsError
}
