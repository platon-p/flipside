package model

type CardSet struct {
	Id      int
	Title   string
	Slug    string
    OwnerId int `db:"owner_id"`
}
