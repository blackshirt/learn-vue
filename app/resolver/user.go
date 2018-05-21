package resolver

import (
	"log"

	"../models"
	"github.com/graphql-go/graphql"
)

type Resolver struct {
	store models.Store
}

func (env *Resolver) AllUsers(params graphql.ResolveParams) (interface{}, error) {
	users, err := env.Store.AllUsers()
	if err != nil {
		log.Fatalf("resolver users error:", err)
		return nil, err
	}
	return users, nil
}
