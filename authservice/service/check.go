package service

import (
	"errors"

	"github.com/platon-p/flipside/authservice/repository"
)

var (
	EmailIncorrectFormatError = errors.New("incorrect format of email")
	EmailExistsError          = errors.New("email already exists")

	NicknameIncorrectFormatError = errors.New("incorrect format of nickname")
	NicknameExistsError          = errors.New("nickname already exists")
)

type CheckService struct {
	userRepository repository.UserRepository
}

func NewCheckService(userRepository repository.UserRepository) *CheckService {
	return &CheckService{
		userRepository: userRepository,
	}
}

// CheckEmail checks if the email exist for a user
// TODO: check for correctness
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

// CheckNickname checks if the nickname exist for a user
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
