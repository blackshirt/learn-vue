package models

import (
	"database/sql"
	"encoding/json"
	"log"
	"reflect"

	"golang.org/x/crypto/bcrypt"
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
	Uid       int    `json: "uid"`
	Username  string `json: "username"`
	HashedPwd string `json: "hashedPwd"`
}

func (repo Repo) GetUserByID(uid int) (*User, error) {
	user := &User{}
	// query db and scan the result
	// use the pointer address
	err := repo.db.QueryRow("select uid, username, hashedPwd from users where uid=?", uid).Scan(
		&user.Uid,
		&user.Username,
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

func (repo Repo) AllUsers() ([]*User, error) {
	rows, err := repo.db.Query("select uid, username, hashedPwd from users")
	if err != nil {
		log.Fatalf("error in select from users:", err.Error())
	}
	defer rows.Close()
	var users []*User
	for rows.Next() {
		user := new(User)
		err := rows.Scan(&user.Uid, &user.Username, &user.HashedPwd)
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

func (repo Repo) CreateUser(user *User) error {

	// store using bcrypt package
	hashedpwd, err := bcrypt.GenerateFromPassword([]byte(user.HashedPwd), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error hashing")
		return err
	}
	// string
	user.HashedPwd = string(hashedpwd)
	// prepare query
	query := `insert into users(username,hashedPwd) values(?,?)`
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Fatalf("error in prepare of create user:", err)
		return err
	}

	_, er := stmt.Exec(user.Username, user.HashedPwd)
	if er != nil {
		log.Fatalf("Error in Exec:", er)
		return er
	}
	//id, err := result.LastInsertId()
	//if err != nil {
	//	log.Fatalf("Error in result.LastinsertId")
	//}

	return nil
}

// Get roles from user with id
func (repo Repo) GetUserRoles(id int) ([]*Role, error) {
	var roles []*Role
	// rows just select the required field to construct Role in rows.Scan
	rows, err := repo.db.Query(`select r.rid, r.name, r.description from roles r 
			join user_role ur on ur.role_id = r.rid where ur.user_id=?`, id)
	// check error
	if err != nil {
		log.Fatalf("error in select from user_role:", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		role := new(Role)
		err := rows.Scan(&role.Rid, &role.Name, &role.Description)
		if err != nil {
			log.Fatalf("Error in rows scans of users role:", err)
		}
		roles = append(roles, role)

	}
	// check error after Next() loop
	if err = rows.Err(); err != nil {
		rows.Close()
	}
	return roles, nil
}

type Role struct {
	Rid         int    `json:"rid"`
	Name        string `json: "name"`
	Description string `json: "description"`
}

type Roles map[string]Role

type Resource struct {
	Resid  int    `json: "resid"`
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
