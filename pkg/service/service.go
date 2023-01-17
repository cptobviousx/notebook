package service

import (
	"github.com/cptobviousx/notebook"
	"github.com/cptobviousx/notebook/pkg/repository"
)

type Authorization interface {
	CreateUser(user notebook.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type NoteBookList interface {
	Create(userId int, list notebook.NoteBookList) (int, error)
	GetAll(userId int) ([]notebook.NoteBookList, error)
	GetById(userId, listId int) (notebook.NoteBookList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input notebook.UpdateListInput) error
}

type NoteBookItem interface {
	Create(userId, listId int, item notebook.NoteBookItem) (int, error)
	GetAll(userId, listId int) ([]notebook.NoteBookItem, error)
	GetById(userId, itemId int) (notebook.NoteBookItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input notebook.UpdateItemInput) error
}

type Service struct {
	Authorization
	NoteBookList
	NoteBookItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		NoteBookList:  NewNoteBookListService(repos.NoteBookList),
		NoteBookItem:  NewNoteBookItemService(repos.NoteBookItem, repos.NoteBookList),
	}
}
