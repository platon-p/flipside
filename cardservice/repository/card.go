package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/platon-p/flipside/cardservice/model"
)

var (
	cardsTable = "cards"
)

type CardRepository interface {
	CreateCards(cards []model.Card) ([]model.Card, error)
	GetCards(slug string) ([]model.Card, error)
	UpdateCards(cards []model.Card) ([]model.Card, error)
	DeleteCards(slug string, positions []int) error
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
        VALUES (:question, :answer, :position, :card_set_id) RETURNING *`,
		cardsTable,
	)
	rows, err := r.db.NamedQuery(query, cards)
	if err != nil {
		return nil, err
	}
	var res []model.Card
	err = rows.StructScan(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *CardRepositoryImpl) GetCards(cardSetSlug string) ([]model.Card, error) {
	query := fmt.Sprintf(
		`UPDATE %v
        SET question = :question, answer = :answer,
        position = :position, card_set_id = :card_set_id
        WHERE position = :position AND card_set_id = :card_set_id
        RETURNING *`,
		cardsTable,
	)
}

func (r *CardRepositoryImpl) UpdateCards(cards []model.Card) ([]model.Card, error) {
	query := fmt.Sprintf(
		`UPDATE %v
        SET question = :question, answer = :answer,
        position = :position, card_set_id = :card_set_id
        WHERE position = :position AND card_set_id = :card_set_id
        RETURNING *`,
		cardsTable,
	)
	rows, err := r.db.NamedQuery(query, cards)
	if err != nil {
		return nil, err
	}
	var res []model.Card
	err = rows.StructScan(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *CardRepositoryImpl) DeleteCards(slug string, positions []int) error {
	query := fmt.Sprintf(
		`DELETE c FROM %v c
        INNER JOIN %v s ON s.id = c.card_set_id
        WHERE s.slug = $1 AND c.position in $2`,
		cardsTable,
        cardSetsTable,
	)
	_, err := r.db.Queryx(query, slug, positions)
	return err
}
