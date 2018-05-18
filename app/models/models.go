package models

type User struct {
	UID       int    `json: "uid"`
	Username  string `json: "username"`
	Firstname string `json: "firstname"`
	Lastname  string `json:"lastname"`
	HashedPwd string `json: "hashedPwd"`
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
