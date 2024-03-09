package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/platon-p/flipside/cardservice/model"
)

var (
	cardsTable         = "cards"
	positionConstraint = "unique_position_per_set"

	ErrCardWithThisPositionExists = errors.New("Card with this position already exists")
	ErrCardNotFound               = errors.New("Card not found")
)

type CardRepository interface {
	CreateCards(cards []model.Card) ([]model.Card, error)
	GetCards(slug string) ([]model.Card, error)
	GetCardsByCardSet(cardSetId int) ([]model.Card, error)
	GetCard(cardId int) (*model.Card, error)
	UpdateCards(cards []model.Card) ([]model.Card, error)
	DeleteCards(cardSetId int, positions []int) error
}

type CardRepositoryImpl struct {
	db *sqlx.DB
}

func NewCardRepositoryImpl(db *sqlx.DB) *CardRepositoryImpl {
	return &CardRepositoryImpl{
		db: db,
	}
}

func (r *CardRepositoryImpl) CreateCards(cards []model.Card) ([]model.Card, error) {
	query := fmt.Sprintf(
		`INSERT INTO %v(question, answer, position, card_set_id)
        VALUES ($1, $2, $3, $4) RETURNING *;`,
		cardsTable,
	)
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}
	stmt, err := tx.Preparex(query)
	if err != nil {
		return nil, err
	}
	res := make([]model.Card, len(cards))
	for i, v := range cards {
		err := stmt.QueryRowx(v.Question, v.Answer, v.Position, v.CardSetId).StructScan(&res[i])
		if err != nil {
			switch e := err.(type) {
			case *pq.Error:
				if e.Code == "23505" && e.Constraint == positionConstraint {
					return nil, ErrCardWithThisPositionExists
				}
				fmt.Printf("%+v\n", *e)
			default:
				fmt.Println(err)
			}
			if err := tx.Tx.Rollback(); err != nil {
				return nil, err
			}
			return nil, err
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *CardRepositoryImpl) GetCardsByCardSet(cardSetId int) ([]model.Card, error) {
	query := fmt.Sprintf(
		`SELECT * FROM %v 
        WHERE card_set_id = $1
        ORDER BY position`,
		cardsTable,
	)
	var res []model.Card
	rows, err := r.db.Queryx(query, cardSetId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var cur model.Card
		if err := rows.StructScan(&cur); err != nil {
			return nil, err
		}
		res = append(res, cur)
	}
	return res, nil
}

func (r *CardRepositoryImpl) GetCards(cardSetSlug string) ([]model.Card, error) {
	cardSet, err := r.getCardSet(cardSetSlug)
	if err != nil {
		return nil, err
	}
	return r.GetCardsByCardSet(cardSet.Id)
}

func (r *CardRepositoryImpl) GetCard(cardId int) (*model.Card, error) {
    query := fmt.Sprintf(
        `SELECT * FROM %v WHERE id = $1`,
        cardsTable,
    )
    var found model.Card
    err := r.db.QueryRowx(query, cardId).StructScan(&found)
    if errors.Is(err, sql.ErrNoRows) {
        return nil, ErrCardNotFound
    }
    if err != nil {
        return nil, err
    }
    return &found, nil
    
}

func (r *CardRepositoryImpl) UpdateCards(cards []model.Card) ([]model.Card, error) {
	query := fmt.Sprintf(
		`UPDATE %v
        SET question = $1, answer = $2
        WHERE position = $3 AND card_set_id = $4
        RETURNING *`,
		cardsTable,
	)
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}
	stmt, err := tx.Preparex(query)
	if err != nil {
		return nil, err
	}
	res := make([]model.Card, len(cards))
	for i, v := range cards {
		err := stmt.QueryRowx(v.Question, v.Answer, v.Position, v.CardSetId).StructScan(&res[i])
		if err != nil {
			if err := tx.Rollback(); err != nil {
				return nil, err
			}
			if errors.Is(err, sql.ErrNoRows) {
				return nil, ErrCardNotFound
			}
			fmt.Println(err)
			return nil, err
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *CardRepositoryImpl) DeleteCards(cardSetId int, positions []int) error {
	query := fmt.Sprintf(
		`DELETE FROM %v WHERE card_set_id = $1 AND position = $2`,
		cardsTable,
	)
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	stmt, err := tx.Preparex(query)
	if err != nil {
		return err
	}
	for _, v := range positions {
		res, err := stmt.Exec(cardSetId, v)
		if err != nil {
			if err := tx.Rollback(); err != nil {
				return err
			}
			return err
		}
		c, err := res.RowsAffected()
		if c != 0 {
			continue
		}
		if err := tx.Rollback(); err != nil {
			return err
		}
		if err != nil {
			return err
		}
		return ErrCardNotFound
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return err
}

func (r *CardRepositoryImpl) getCardSet(slug string) (*model.CardSet, error) {
	var found model.CardSet
	query := fmt.Sprintf(`SELECT * FROM %v WHERE slug = $1`, cardSetsTable)
	err := r.db.QueryRowx(query, slug).StructScan(&found)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrCardSetNotFound
	}
	if err != nil {
		return nil, err
	}
	return &found, nil
}
