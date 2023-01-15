package service

import (
	"github.com/cptobviousx/notebook"
	"github.com/cptobviousx/notebook/pkg/repository"
)

type NoteBookListService struct {
	repo repository.NoteBookList
}

func NewNoteBookListService(repo repository.NoteBookList) *NoteBookListService {
	return &NoteBookListService{repo: repo}
}

func (s *NoteBookListService) Create(userId int, list notebook.NoteBookList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *NoteBookListService) GetAll(userId int) ([]notebook.NoteBookList, error) {
	return s.repo.GetAll(userId)
}

func (s *NoteBookListService) GetById(userId, listId int) (notebook.NoteBookList, error) {
	return s.repo.GetById(userId, listId)
}
