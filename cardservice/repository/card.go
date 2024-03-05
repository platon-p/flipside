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
	GetCards(cardSetSlug string) ([]model.Card, error)
	UpdateCards(cards []model.Card) ([]model.Card, error)
	DeleteCards(cards []model.Card) error
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

func (r *CardRepositoryImpl) DeleteCards(cards []model.Card) error {
	query := fmt.Sprintf(
		`DELETE FROM %v
        WHERE position = :position AND card_set_id = :card_set_id`,
		cardsTable,
	)
	_, err := r.db.NamedQuery(query, cards)
	return err
}
