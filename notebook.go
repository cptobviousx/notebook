package notebook

import (
	"fmt"
)

type NoteBookList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type NoteBookItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type NoteBookItemsList struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return fmt.Errorf("update structure requires a Title and a Description (Title: %d, Description: %d)", i.Title, i.Description)
	}

	return nil
}

func (i UpdateItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return fmt.Errorf("update structure requires data (Title: %d, Description: %d, Done: %d)", i.Title, i.Description, i.Done)
	}

	return nil
}
