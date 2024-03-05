package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/platon-p/flipside/cardservice/model"
)

var (
	cardSetsTable = "card_sets"
)

type CardSetRepository interface {
	CreateCardSet(cardSet *model.CardSet) (*model.CardSet, error)
	GetCardSet(slug string) (*model.CardSet, error)
	UpdateCardSet(slug string, cardSet *model.CardSet) (*model.CardSet, error)
	DeleteCardSet(slug string) error
}

type CardSetRepositoryImpl struct {
	db *sqlx.DB
}

func NewCardSetRepositoryImpl(db *sqlx.DB) *CardSetRepositoryImpl {
	return &CardSetRepositoryImpl{
		db: db,
	}
}

func (r *CardSetRepositoryImpl) CreateCardSet(cardSet *model.CardSet) (*model.CardSet, error) {
	query := fmt.Sprintf(
		`INSERT INTO %v(title, slug, owner_id)
        VALUES($1, $2, $3)
        RETURNING *`,
		cardSetsTable,
	)
	var newEntity model.CardSet
	err := r.db.QueryRowx(query, cardSet.Title, cardSet.Title, cardSet.OwnerId).StructScan(&newEntity)
	if err != nil {
		return nil, err
	}
	return &newEntity, nil
}

func (r *CardSetRepositoryImpl) GetCardSet(slug string) (*model.CardSet, error) {
	var found model.CardSet
	query := fmt.Sprintf(`SELECT * FROM %v WHERE slug = $1`, cardSetsTable)
	err := r.db.QueryRowx(query, slug).StructScan(&found)
	if err != nil {
		return nil, err
	}
	return &found, nil
}

func (r *CardSetRepositoryImpl) UpdateCardSet(slug string, card *model.CardSet) (*model.CardSet, error) {
	var updated model.CardSet
	query := fmt.Sprintf(
		`UPDATE %v 
        SET title = $1, slug = $2, owner_id = $3 
        WHERE slug = $2
        RETURNING *`,
		cardSetsTable,
	)
    err := r.db.QueryRowx(query, card.Title, card.Slug, card.OwnerId).StructScan(&updated)
    if err != nil {
        return nil, err
    }
    return &updated, nil
}

func (r *CardSetRepositoryImpl) DeleteCardSet(slug string) error {
    query := fmt.Sprintf(`DELETE FROM %v WHERE slug = $1`, cardSetsTable)
    _, err := r.db.Exec(query, slug)
    return err
}
