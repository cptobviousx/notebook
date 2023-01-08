package repository

import (
	"github.com/cptobviousx/notebook"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user notebook.User) (int, error)
	GetUser(username, password string) (notebook.User, error)
}

type NoteBookList interface {
}

type NoteBookItem interface {
}

type Repository struct {
	Authorization
	NoteBookList
	NoteBookItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
