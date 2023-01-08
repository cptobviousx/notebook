package repository

import (
	"fmt"

	"github.com/cptobviousx/notebook"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user notebook.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT  INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (notebook.User, error) {
	var user notebook.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 and password_hash = $2", usersTable)

	err := r.db.Get(&user, query, username, password)

	return user, err
}
