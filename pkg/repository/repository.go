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
	Create(userId int, list notebook.NoteBookList) (int, error)
	GetAll(userId int) ([]notebook.NoteBookList, error)
	GetById(userId, listId int) (notebook.NoteBookList, error)
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
		NoteBookList:  NewNoteBookPostgres(db),
	}
}
