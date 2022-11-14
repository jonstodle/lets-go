package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func Connect(dbPath string) (Database, error) {
	if db, err := sqlx.Connect("sqlite3", dbPath); err != nil {
		return nil, err
	} else {
		return db, nil
	}
}

type Database interface {
	Get(dest any, query string, args ...any) error
	Select(dest any, query string, args ...any) error
}
