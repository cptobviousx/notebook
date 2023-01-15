package repository

import (
	"fmt"

	"github.com/cptobviousx/notebook"
	"github.com/jmoiron/sqlx"
)

type NoteBookListPostgres struct {
	db *sqlx.DB
}

func NewNoteBookPostgres(db *sqlx.DB) *NoteBookListPostgres {
	return &NoteBookListPostgres{db: db}
}

func (r *NoteBookListPostgres) Create(userId int, list notebook.NoteBookList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", noteBookListTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *NoteBookListPostgres) GetAll(userId int) ([]notebook.NoteBookList, error) {
	var lists []notebook.NoteBookList
	query := fmt.Sprintf("SELECT t1.id, t1.title, t1.description FROM %s t1 INNER JOIN %s t2 ON t1.id = t2.list_id WHERE t2.user_id = $1", noteBookListTable, usersListsTable)

	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *NoteBookListPostgres) GetById(userId, listId int) (notebook.NoteBookList, error) {
	var list notebook.NoteBookList

	query := fmt.Sprintf("SELECT t1.id, t1.title, t1.description FROM %s t1 INNER JOIN %s t2 ON t1.id = t2.list_id WHERE t2.user_id = $1 AND t2.list_id = $2", noteBookListTable, usersListsTable)

	err := r.db.Get(&list, query, userId, listId)

	return list, err
}
