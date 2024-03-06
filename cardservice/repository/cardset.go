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
	cardSetsTable = "card_sets"
    cardSetsSlugConstraint = `card_sets_slug_key`

	ErrCardSetNotFound = errors.New("Card Set not found")
    ErrCardSetSlugAlreadyExists = errors.New("Slug already exists")
)

type CardSetRepository interface {
	CreateCardSet(cardSet *model.CardSet) (*model.CardSet, error)
	GetCardSet(slug string) (*model.CardSet, error)
	UpdateCardSet(id int, cardSet *model.CardSet) (*model.CardSet, error)
	DeleteCardSet(ownerId int, cardSetId int) error
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
	err := r.db.QueryRowx(query, cardSet.Title, cardSet.Slug, cardSet.OwnerId).StructScan(&newEntity)
    switch e := err.(type) {
    case *pq.Error:
        if e.Constraint == cardSetsSlugConstraint {
            return nil, ErrCardSetSlugAlreadyExists
        }
        fmt.Println(*e)
    }
	if err != nil {
		return nil, err
	}
	return &newEntity, nil
}

func (r *CardSetRepositoryImpl) GetCardSet(slug string) (*model.CardSet, error) {
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

func (r *CardSetRepositoryImpl) UpdateCardSet(id int, card *model.CardSet) (*model.CardSet, error) {
	var updated model.CardSet
	query := fmt.Sprintf(
		`UPDATE %v 
        SET title = $1, slug = $2
        WHERE id = $4, owner_id = $3
        RETURNING *`,
		cardSetsTable,
	)
	err := r.db.QueryRowx(query, card.Title, card.Slug, card.OwnerId, id).StructScan(&updated)
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func (r *CardSetRepositoryImpl) DeleteCardSet(ownerId int, id int) error {
	query := fmt.Sprintf(`DELETE FROM %v WHERE id = $1 AND owner_id = $2`, cardSetsTable)
	_, err := r.db.Exec(query, id, ownerId)
	return err
}
