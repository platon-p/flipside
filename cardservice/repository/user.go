package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/platon-p/flipside/cardservice/model"
	"golang.org/x/text/number"
)

var (
    usersTable = "users"

    ErrProfileNotFound = errors.New("Profile not found")
)

type UserRepository interface {
    GetProfile(nickname string) (*model.Profile, error)
    GetCardSets(userId int) ([]model.CardSet, error)
}

type UserRepositoryImpl struct {
    db *sqlx.DB
}

func (r *UserRepositoryImpl) GetProfile(nickname string) (*model.Profile, error) {
    query := fmt.Sprintf(
        `SELECT id, name, nickname FROM %v
        WHERE nickname = $1`,
        usersTable,
    )
    var found model.Profile
    err := r.db.QueryRowx(query, nickname).StructScan(&found)
    if errors.Is(err, sql.ErrNoRows) {
        return nil, ErrProfileNotFound
    }
    if err != nil {
        return nil, err
    }
    return &found, nil
}

func (r *UserRepositoryImpl) GetCardSets(userId int) ([]model.CardSet, error) {
    query := fmt.Sprintf(
        `SELECT * FROM %v
        WHERE owner_id = $1`,
        cardSetsTable,
    )
    res := make([]model.CardSet, 0)
    rows, err := r.db.Queryx(query, userId)
    for rows.Next() {
        var row model.CardSet
        if err := rows.StructScan(&row); err != nil {
            return nil, err
        }
        res = append(res, row)
    }
    if err != nil {
        return nil, err
    }
    return res, nil
}
