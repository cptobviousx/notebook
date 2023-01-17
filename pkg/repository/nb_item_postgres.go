package repository

import (
	"fmt"
	"strings"

	"github.com/cptobviousx/notebook"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type NoteBookItemPostgres struct {
	db *sqlx.DB
}

func NewNoteBookItemPostgres(db *sqlx.DB) *NoteBookItemPostgres {
	return &NoteBookItemPostgres{db: db}
}

func (r *NoteBookItemPostgres) Create(listId int, item notebook.NoteBookItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING ID", notebookItemsTable)

	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listsItemsTable)

	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *NoteBookItemPostgres) GetAll(userId, listId int) ([]notebook.NoteBookItem, error) {
	var items []notebook.NoteBookItem

	query := fmt.Sprintf("SELECT t1.id, t1.title, t1.description, t1.done FROM %s t1 INNER JOIN %s t2 on t1.id = t2.item_id INNER JOIN %s t3 on t2.list_id = t3.list_id WHERE t2.list_id = $1 AND t3.user_id = $2", notebookItemsTable, listsItemsTable, usersListsTable)

	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *NoteBookItemPostgres) GetById(userId, itemId int) (notebook.NoteBookItem, error) {
	var item notebook.NoteBookItem

	query := fmt.Sprintf("SELECT t1.id, t1.title, t1.description, t1.done FROM %s t1 INNER JOIN %s t2 on t1.id = t2.item_id INNER JOIN %s t3 on t2.list_id = t3.list_id WHERE t1.id = $1 AND t3.user_id = $2", notebookItemsTable, listsItemsTable, usersListsTable)

	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *NoteBookItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s t1 USING %s t2, %s t3 WHERE t1.id = t2.item_id AND t2.list_id = t3.list_id AND t3.user_id = $1 AND t1.id = $2", notebookItemsTable, listsItemsTable, usersListsTable)

	_, err := r.db.Exec(query, userId, itemId)

	return err
}

func (r *NoteBookItemPostgres) Update(userId, itemId int, input notebook.UpdateItemInput) error {
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

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s ni SET %s FROM %s li, %s ul WHERE ni.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ni.id = $%d",
		notebookItemsTable, setQuery, listsItemsTable, usersListsTable, argId, argId+1)
	args = append(args, userId, itemId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
