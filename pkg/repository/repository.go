package repository

type Authorization interface {
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

func NewRepository() *Repository {
	return &Repository{}
}
