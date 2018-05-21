package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Datastore interface {
	AllUsers() (interface{}, error)
}

type SQLResolver struct {
	DB *sql.DB
}

func NewDB(dataSourceName string) (*SQLResolver, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatalf("DB open error", err.Error())
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Ping error", err.Error())
		return nil, err
	}
	return &SQLResolver{db}, nil
}
