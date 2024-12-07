package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type database struct {
	*sql.DB
}

func New() (*database, error) {
	db, err := sql.Open("sqlite3", "./mcp.db")
	if err != nil {
		return nil, err
	}
	return &database{DB: db}, nil
}
