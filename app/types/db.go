package types

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SQLResolver struct {
	DB *sql.DB
}

func Open(dataSourceName string) (*SQLResolver, error) {
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

// initiliaze with test database, remember to use right path of db from the root
// of the main package. .. even at the same path with this package !!!
var db, _ = Open("./app/types/test.db")

func (sqh *SQLResolver) GetByID(id int) (*User, error) {
	user := &User{}
	// query db and scan the result
	// use the pointer address
	err := sqh.DB.QueryRow("select id, username, firstname, lastname from user where id=?", id).Scan(
		&user.ID,
		&user.Username,
		&user.Firstname,
		&user.Lastname)
	// check error
	if err != nil {
		// check if QueryRow may result in empty result, no row in result set error
		if err == sql.ErrNoRows {
			// just return the empty user
			return user, nil
		}
		log.Fatalf("GetByID query error:", err.Error())
		return nil, err
	}
	return user, nil
}

func (sqh *SQLResolver) Users() ([]*User, error) {
	var users []*User
	rows, err := sqh.DB.Query("select id, username, firstname, lastname from user")
	// check error
	if err != nil {
		log.Fatalf("error in select from user:", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		user := new(User)
		err := rows.Scan(&user.ID, &user.Username, &user.Firstname, &user.Lastname)
		if err != nil {
			log.Fatalf("Error in rows scans of users:", err)
		}
		users = append(users, user)

	}
	// check error after Next() loop
	if err = rows.Err(); err != nil {
		rows.Close()
	}
	return users, nil
}
