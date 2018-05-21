package types

import (
	"log"

	"github.com/graphql-go/graphql"
	_ "github.com/mattn/go-sqlite3"

	"../models"
)

var db, _ = models.NewDB("./app/types/test.db")

var userObjectType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "Graphql object type for User model",
	Fields: graphql.Fields{
		"uid": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "Id of the user",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				if user, ok := params.Source.(*models.User); ok {
					return user.Uid, nil
				}
				return nil, nil
			},
		},
		"username": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Username of User model",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				if user, ok := params.Source.(*models.User); ok {
					return user.Username, nil
				}
				return nil, nil

			},
		},
		"passwd": &graphql.Field{
			Type:        graphql.String,
			Description: "Passwd",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				if user, ok := params.Source.(*models.User); ok {
					// remember, hashedPwd was Nullablestring type from models,
					// pick up the String, even its NULL
					return user.HashedPwd, nil
				}
				return nil, nil
			},
		},
		"roles": &graphql.Field{
			Type:        graphql.NewList(roleObjectType),
			Description: "roles of the user",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// var roles []Role
				userID := params.Source.(*models.User).Uid
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
		"rid": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.ID),
			Description: "Id of the Role",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				if role, ok := params.Source.(*models.Role); ok {
					return role.Rid, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "name field of the role",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				if role, ok := params.Source.(*models.Role); ok {
					return role.Name, nil
				}
				return nil, nil
			},
		},
		"description": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "description field of the role",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				if role, ok := params.Source.(*models.Role); ok {
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
				"uid": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				uid, _ := params.Args["uid"].(int)
				// db was database handler in db files in the same package
				user, err := db.GetUserByID(uid)
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
				users, err := db.AllUsers()
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
				"passwd": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "password input to create user mutation",
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				user := &models.User{
					Username:  params.Args["username"].(string),
					HashedPwd: params.Args["passwd"].(string),
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
