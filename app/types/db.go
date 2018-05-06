package types

import "database/sql"

type SQLResolver struct {
	Conn *sql.DB
}

func NewDB(driver, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open(driver, dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func NewSQLHandler(conn *sql.DB) *SQLResolver {
	return &SQLResolver{Conn: conn}
}

func (sqh *SQLResolver) GetByID(id int) (*User, error) {
	user := &User{}
	// query db and scan the result
	err := sqh.Conn.QueryRow("select id, username, firstname, lastname from user where id=?", id).Scan(
		user.ID,
		user.Username,
		user.Firstname,
		user.Lastname)
	// check error
	if err != nil {
		panic(err)
	}
	return user, nil
}

func (sqh *SQLResolver) Users() ([]*User, error) {
	var users []*User
	rows, err := sqh.Conn.Query("select * from user")
	// check error
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		user := new(User)
		err := rows.Scan(&user.ID, &user.Username, &user.Firstname, &user.Lastname)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	return users, nil
}
