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
	Delete(userId, listId int) error
	Update(userId, listId int, input notebook.UpdateListInput) error
}

type NoteBookItem interface {
	Create(listId int, item notebook.NoteBookItem) (int, error)
	GetAll(userId, listId int) ([]notebook.NoteBookItem, error)
	GetById(userId, itemId int) (notebook.NoteBookItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input notebook.UpdateItemInput) error
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
		NoteBookItem:  NewNoteBookItemPostgres(db),
	}
}
