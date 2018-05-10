package types

import (
	"log"

	"github.com/graphql-go/graphql"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID        int    `json: "id"`
	Username  string `json: "username"`
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

type Role struct {
	ID          int    `json: "id"`
	Name        string `json: "name"`
	Description string `json: "description"`
}

var userObjectType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "Graphql object type for User model",
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
		"roles": &graphql.Field{
			Type:        graphql.NewList(roleObjectType),
			Description: "roles of the user",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// var roles []Role
				userID := params.Source.(*User).ID
				roles, err := db.GetUserRoles(userID)
				if err != nil {
					log.Fatalf("Error get user roles", err)
					return nil, err
				}

				return roles, nil
			},
		},
	},
})

var roleObjectType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Role",
	Description: "Graphql role object type",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.ID),
			Description: "Id of the Role",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				if role, ok := params.Source.(*Role); ok {
					return role.ID, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "name field of the role",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				if role, ok := params.Source.(*Role); ok {
					return role.Name, nil
				}
				return nil, nil
			},
		},
		"description": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "description field of the role",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				if role, ok := params.Source.(*Role); ok {
					return role.Description, nil
				}
				return nil, nil
			},
		},
	},
})

// Query type object
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type:        userObjectType,
			Description: "Get an user",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, _ := params.Args["id"].(int)
				// db was database handler in db files in the same package
				user, err := db.GetByID(id)
				if err != nil {
					log.Fatalf("Error in resolver getbyID", err)
					return nil, err
				}
				return user, nil
			},
		},

		// multiple set
		"users": &graphql.Field{
			Type:        graphql.NewList(userObjectType),
			Description: "List of user",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				users, err := db.Users()
				if err != nil {
					log.Fatalf("resolver users error:", err)
					return nil, err
				}
				return users, nil
			},
		},
	},
})

// Mutations type object
var MutationsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutations",
	Fields: graphql.Fields{
		"createUserMutation": &graphql.Field{
			Type:        userObjectType,
			Description: "create user mutation",
			Args: graphql.FieldConfigArgument{
				"username": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "username input to create user mutation",
				},
				"firstname": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"lastname": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				user := &User{
					Username:  params.Args["username"].(string),
					Firstname: params.Args["firstname"].(string),
					Lastname:  params.Args["lastname"].(string),
				}
				// add db logic here
				err := db.CreateUser(user)
				if err != nil {
					return nil, err
				}
				return user, nil
			},
		},
	},
})
