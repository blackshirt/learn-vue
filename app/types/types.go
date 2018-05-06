package types

import (
	"github.com/graphql-go/graphql"
)

// var db *sql.DB

type User struct {
	ID        int    `json: "id"`
	Username  string `json: "username"`
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

var db, _ = NewDB("sqlite3", "./test.db")

var resolver = NewSQLHandler(db)

var UserTypeObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "Graphql type object for User model",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "Id of the user",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				if user, ok := params.Source.(*User); ok {
					return user.ID, nil
				}
				return nil, nil
			},
		},
		"username": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Username of User model",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				if user, ok := params.Source.(*User); ok {
					return user.Username, nil
				}
				return nil, nil

			},
		},
		"firstname": &graphql.Field{
			Type:        graphql.String,
			Description: "Firstname",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				if user, ok := params.Source.(*User); ok {
					return user.Firstname, nil
				}
				return nil, nil
			},
		},
		"lastname": &graphql.Field{
			Type:        graphql.String,
			Description: "Lastname",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				if user, ok := params.Source.(*User); ok {
					return user.Lastname, nil
				}
				return nil, nil
			},
		},
	},
})

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type:        UserTypeObject,
			Description: "Get an user",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, _ := params.Args["id"].(int)
				user, err := resolver.GetByID(id)
				if err != nil {
					return nil, err
				}
				return user, nil
			},
		},

		// multiple set
		"users": &graphql.Field{
			Type:        graphql.NewList(UserTypeObject),
			Description: "List of user",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				users, err := resolver.Users()
				if err != nil {
					return nil, err
				}
				return users, nil
			},
		},
	},
})
