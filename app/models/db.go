package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Store interface {
	AllUsers() (interface{}, error)
}

type Repo struct {
	db *sql.DB
}

func NewDB(dataSourceName string) (*Repo, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatalf("DB open error", err.Error())
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Ping error", err.Error())
		return nil, err
	}
	return &Repo{db}, nil
}
