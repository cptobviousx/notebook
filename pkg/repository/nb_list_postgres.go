package repository

import (
	"fmt"
	"strings"

	"github.com/cptobviousx/notebook"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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

func (r *NoteBookListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s t1 USING %s t2 WHERE t1.id = t2.list_id AND t2.user_id = $1 AND t2.list_id = $2", noteBookListTable, usersListsTable)

	_, err := r.db.Exec(query, userId, listId)

	return err
}

func (r *NoteBookListPostgres) Update(userId, listId int, input notebook.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	// title=$1
	// description=$1
	// title=$1, description=$2
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		noteBookListTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, listId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
