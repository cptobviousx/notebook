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
}

type NoteBookItem interface {
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
	}
}
