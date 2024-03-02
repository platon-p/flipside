package repository

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/platon-p/flashside/authservice/model"
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
    db sqlx.DB
}

func (r *RefreshTokenRepositoryPostgres) Create(userId int, token string, expiresAt time.Time) (*model.RefreshToken, error) {
    var newEntity model.RefreshToken
    query := fmt.Sprintf("INSERT INTO %v(token, user_id, expires_at) VALUES ($1, $2, $3)", refreshTokenTable)
    err := r.db.QueryRowx(query, token, userId, expiresAt).StructScan(&newEntity)
    if err != nil {
        return nil, err
    }
    return &newEntity, nil
}

func (r *RefreshTokenRepositoryPostgres) FindByToken(token string) (*model.RefreshToken, error) {
	var found model.RefreshToken
    query := fmt.Sprintf("SELECT * FROM %v WHERE token = ?", refreshTokenTable)
	err := r.db.QueryRowx(query, token).Scan(&found)
	if err != nil {
		return nil, err
	}
	return &found, nil
}

func (r *RefreshTokenRepositoryPostgres) FindByUser(userId int) (*model.RefreshToken, error) {
	var found model.RefreshToken
    query := fmt.Sprintf("SELECT * FROM %v WHERE user_id = ?", refreshTokenTable)
	err := r.db.QueryRowx(query, userId).Scan(&found)
	if err != nil {
		return nil, err
	}
	return &found, nil
}

func (r *RefreshTokenRepositoryPostgres) Delete(token string) error {
    query := fmt.Sprintf("DELETE FROM %v WHERE token = ?", refreshTokenTable)
    _, err := r.db.Exec(query, token)
    return err
}
