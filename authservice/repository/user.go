package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/platon-p/flipside/authservice/model"
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
	db *sqlx.DB
}

func NewUserRepositoryImpl(db *sqlx.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) Create(user *model.User) (*model.User, error) {
	var newEntity model.User
	query := fmt.Sprintf(
		`INSERT INTO %v(created_at, name, nickname, email, password) 
        VALUES ($1, $2, $3, $4, $5) 
        RETURNING *;`, usersTable)
	err := r.db.
		QueryRowx(query, time.Now(), user.Name, user.Nickname, user.Email, user.Password).
		StructScan(&newEntity)
	if err != nil {
		return nil, err
	}
	return &newEntity, nil
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*model.User, error) {
	var found model.User
	query := fmt.Sprintf("SELECT * FROM %v WHERE email = $1;", usersTable)
	err := r.db.QueryRowx(query, email).StructScan(&found)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &found, nil
}

func (r *UserRepositoryImpl) FindByNickname(nickname string) (*model.User, error) {
	var found model.User
	query := fmt.Sprintf("SELECT * FROM %v WHERE nickname = $1;", usersTable)
	err := r.db.QueryRowx(query, nickname).StructScan(&found)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &found, nil
}
