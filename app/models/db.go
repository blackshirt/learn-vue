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

func (sqh *SQLResolver) CreateUser(user *User) error {
	// var id int64
	query := `insert into users(username, firstname, lastname) values(?,?,?)`
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
	user_role ur, roles r where ur.user_id=? and ur.role_id = r.id`, id)
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
