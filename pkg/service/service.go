package service

import (
	"github.com/cptobviousx/notebook"
	"github.com/cptobviousx/notebook/pkg/repository"
)

type Authorization interface {
	CreateUser(user notebook.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type NoteBookList interface {
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
	}
}
