package model

type Card struct {
	Id        int
	Question  string
	Answer    string
	Position  int
    CardSetId int `db:"card_set_id"`
}
