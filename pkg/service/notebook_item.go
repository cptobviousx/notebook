package service

import (
	"github.com/cptobviousx/notebook"
	"github.com/cptobviousx/notebook/pkg/repository"
)

type NoteBookItemService struct {
	repo     repository.NoteBookItem
	listRepo repository.NoteBookList
}

func NewNoteBookItemService(repo repository.NoteBookItem, listRepo repository.NoteBookList) *NoteBookItemService {
	return &NoteBookItemService{repo: repo, listRepo: listRepo}
}

func (s *NoteBookItemService) Create(userId, listId int, item notebook.NoteBookItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *NoteBookItemService) GetAll(userId, listId int) ([]notebook.NoteBookItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *NoteBookItemService) GetById(userId, itemId int) (notebook.NoteBookItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *NoteBookItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *NoteBookItemService) Update(userId, itemId int, input notebook.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, itemId, input)
}
