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

func (s *NoteBookListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *NoteBookListService) Update(userId, listId int, input notebook.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}
