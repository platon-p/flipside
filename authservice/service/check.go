package service

import "errors"

var (
    EmailIncorrectFormatError = errors.New("Incorrect format of email")
    EmailExistsError = errors.New("Email already exists")

    NicknameIncorrectFormatError = errors.New("Incorect format of nickname")
    NicknameExistsError = errors.New("Nickname already exists")
)

type CheckService struct{}

func (s *CheckService) CheckEmail(email string) error

func (s *CheckService) CheckNickname(nickname string) error
