package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/platon-p/flashside/authservice/model"
)

var (
	usersTable = "users"
)

type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByNickname(nickname string) (*model.User, error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) *UserRepositoryImpl {
    return &UserRepositoryImpl{
        db: db,
    }
}

func (r *UserRepositoryImpl) Create(user *model.User) (*model.User, error) {
	var newEntity model.User
	query := fmt.Sprintf("INSERT INTO %v(created_at, name, nickname, email, password) VALUES ($1, $2, $3, $4, $5)", usersTable)
	err := r.db.
		QueryRow(query, time.Now(), user.Name, user.Nickname, user.Email, user.Password).
		Scan(&newEntity)
    if err != nil {
        return nil, err
    }
    return &newEntity, nil
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*model.User, error) {
	var found model.User
    query := fmt.Sprintf("SELECT * FROM %v WHERE email = ?", usersTable)
	err := r.db.QueryRow(query, email).Scan(&found)
	if err != nil {
		return nil, err
	}
	return &found, nil
}

func (r *UserRepositoryImpl) FindByNickname(nickname string) (*model.User, error) {
	var found model.User
    query := fmt.Sprintf("SELECT * FROM %v WHERE nickname = ?", usersTable)
	err := r.db.QueryRow(query, nickname).Scan(&found)
	if err != nil {
		return nil, err
	}
	return &found, nil
}
