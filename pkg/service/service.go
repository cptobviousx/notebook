package service

import "github.com/cptobviousx/notebook/pkg/repository"

type Authorization interface {
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
	return &Service{}
}
