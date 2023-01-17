package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable         = "users"
	noteBookListTable  = "notebook_lists"
	usersListsTable    = "users_lists"
	notebookItemsTable = "notebooks_items"
	listsItemsTable    = "lists_items"
)

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	DBName   string `json:"DBName"`
	SSLMode  string `json:"SSLMode"`
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host = %s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
