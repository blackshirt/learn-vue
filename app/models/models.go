package models

import (
	"database/sql"
	"encoding/json"
	"log"
	"reflect"
)

type NullableString sql.NullString

// Scan implements the Scanner interface for NullString
func (ns *NullableString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ns = NullableString{s.String, false}
	} else {
		*ns = NullableString{s.String, true}
	}

	return nil
}

// MarshalJSON for NullString
func (ns *NullableString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	} else {
		return json.Marshal(nil)
	}
	// if !ns.Valid {
	// 	return []byte("null"), nil
	// }
	// return json.Marshal(ns.String)
}

// UnmarshalJSON for NullString
func (ns *NullableString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.String)
	ns.Valid = (err == nil)
	return err
}

type User struct {
	Uid       int            `json: "uid"`
	Username  string         `json: "username"`
	Firstname NullableString `json: "firstname"`
	Lastname  NullableString `json: "lastname"`
	HashedPwd NullableString `json: "hashedPwd"`
}

func (sqr *SQLResolver) GetUserByID(uid int) (*User, error) {
	user := &User{}
	// query db and scan the result
	// use the pointer address
	err := sqr.DB.QueryRow("select uid, username, firstname, lastname, hashedPwd from users where uid=?", uid).Scan(
		&user.Uid,
		&user.Username,
		&user.Firstname,
		&user.Lastname,
		&user.HashedPwd,
	)
	// check error
	if err != nil {
		// check if QueryRow may result in empty result, no row in result set error
		if err == sql.ErrNoRows {
			// just return the empty user
			return user, nil
		}
		log.Fatalf("GetUserByID query error:", err.Error())
		return nil, err
	}
	return user, nil
}

func (sqr *SQLResolver) AllUsers() ([]*User, error) {
	rows, err := sqr.DB.Query("select uid, username, firstname, lastname, hashedPwd from users")
	if err != nil {
		log.Fatalf("error in select from users:", err.Error())
	}
	defer rows.Close()
	var users []*User
	for rows.Next() {
		user := new(User)
		err := rows.Scan(&user.Uid, &user.Username, &user.Firstname, &user.Lastname, &user.HashedPwd)
		if err != nil {
			log.Fatalf("Error in rows scans of users:", err.Error())
		}
		users = append(users, user)

	}
	// check error after Next() loop
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

type Role struct {
	Rid         int    `json:"rid"`
	Name        string `json: "name"`
	Description string `json: "description"`
}

type Roles []Role

type Object struct {
	Obid   int    `json: "obid"`
	Name   string `json:"name"`
	Locked bool   `json:"locked"`
}

type Operation struct {
	Opid   int    `json:"opid"`
	Name   string `json:"name"`
	Locked bool   `json:"locked"`
}

type Permission struct {
	Pid         int    `json:"pid"`
	Name        string `json:"name"`
	Description string `json: "description"`
	Operations  []Operation
	Objects     []Object
}
