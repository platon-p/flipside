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
	refreshTokenTable = "refresh_tokens"
)

type RefreshTokenRepository interface {
	Create(userId int, token string, expiresAt time.Time) (*model.RefreshToken, error)
	FindByToken(token string) (*model.RefreshToken, error)
	FindByUser(userId int) (*model.RefreshToken, error)
	Delete(token string) error
}

type RefreshTokenRepositoryPostgres struct {
	db *sqlx.DB
}

func NewRefreshTokenRepositoryPostgres(db *sqlx.DB) *RefreshTokenRepositoryPostgres {
	return &RefreshTokenRepositoryPostgres{
		db: db,
	}
}

func (r *RefreshTokenRepositoryPostgres) Create(userId int, token string, expiresAt time.Time) (*model.RefreshToken, error) {
	qInsert := fmt.Sprintf("INSERT INTO %v(token, user_id, expires_at) VALUES ($1, $2, $3);", refreshTokenTable)
	err := r.db.QueryRow(qInsert, token, userId, expiresAt).Err()
	if err != nil {
		return nil, err
	}
	return r.FindByToken(token)
}

func (r *RefreshTokenRepositoryPostgres) FindByToken(token string) (*model.RefreshToken, error) {
	qSelect := fmt.Sprintf(
		`SELECT t.token, 
        u.id "user.id", u.name "user.name", u.nickname "user.nickname", u.email "user.email", u.password "user.password",
        t.expires_at
        FROM %v u 
        JOIN %v t ON u.id = t.user_id 
        WHERE t.token = $1`, usersTable, refreshTokenTable)
	var newEntity model.RefreshToken
	err := r.db.QueryRowx(qSelect, token).StructScan(&newEntity)
	if err != nil {
		return nil, err
	}
	return &newEntity, nil
}

func (r *RefreshTokenRepositoryPostgres) FindByUser(userId int) (*model.RefreshToken, error) {
	var found model.RefreshToken
	query := fmt.Sprintf(
		`SELECT t.token, 
        u.id "user.id", u.name "user.name", u.nickname "user.nickname", u.email "user.email", u.password "user.password",
        t.expires_at
        FROM %v u 
        JOIN %v t ON u.id = t.user_id 
        WHERE t.user_id = $1`, usersTable, refreshTokenTable)
	err := r.db.QueryRowx(query, userId).StructScan(&found)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &found, nil
}

func (r *RefreshTokenRepositoryPostgres) Delete(token string) error {
	query := fmt.Sprintf("DELETE FROM %v WHERE token = $1;", refreshTokenTable)
	_, err := r.db.Exec(query, token)
	return err
}
