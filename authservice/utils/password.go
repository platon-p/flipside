package utils

import "golang.org/x/crypto/bcrypt"

type PasswordUtility struct {}

func NewPasswordUtility() *PasswordUtility {
    return &PasswordUtility{}
}

func (u *PasswordUtility) GetPasswordHash(password string) (*string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }
    hashString := string(hash)
    return &hashString, nil
}

func (u *PasswordUtility) CheckPasswordHash(hashedPassword, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err != nil
}
