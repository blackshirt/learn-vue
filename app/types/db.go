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

func (sqh *SQLResolver) CreateUser(user *User) error {
	// var id int64
	query := `insert into user(username, firstname, lastname) values(?,?,?)`
	stmt, err := sqh.DB.Prepare(query)
	if err != nil {
		log.Fatalf("error in prepare of create user:", err)
		return err
	}
	_, er := stmt.Exec(user.Username, user.Firstname, user.Lastname)
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

// Get roles from user
func (sqh *SQLResolver) GetUserRoles(id int) ([]*Role, error) {
	var roles []*Role
	// rows just select the required field to construct Role in rows.Scan
	rows, err := sqh.DB.Query(`select r.id, r.name, r.description from 
	user_role ur, role r where ur.user=? and ur.role = r.id`, id)
	// check error
	if err != nil {
		log.Fatalf("error in select from user_role:", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		role := new(Role)
		err := rows.Scan(&role.ID, &role.Name, &role.Description)
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
