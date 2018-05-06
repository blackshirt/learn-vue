package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	"./app/types"
)

// var db *sql.DB

func main() {
	//db, err := sql.Open("sqlite3", "./test.db")
	//if err != nil {
	//	panic(err)
	//}
	// graphql.Fields map[string]*Field
	// var rootFields = graphql.Fields{
	//	"users": &graphql.Field{
	//		Type: graphql.String,
	//		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
	//			return "paijo", nil
	//		},
	//	},
	// }
	// type ObjectConfig struct {
	// Name        string      `json:"name"`
	// Interfaces  interface{} `json:"interfaces"`
	// Fields      interface{} `json:"fields"`
	// IsTypeOf    IsTypeOfFn  `json:"isTypeOf"`
	// Description string      `json:"description"`
	// }
	// var rootObjectConfig = graphql.ObjectConfig{
	//	Name:   "RootQuery",
	//	Fields: rootFields,
	// }

	// var rootQueryType = graphql.NewObject(rootObjectConfig)
	// type SchemaConfig struct {
	// Query        *Object
	// Mutation     *Object
	// Subscription *Object
	// Types        []Type
	// Directives   []*Directive
	//}
	var schemaConfig = graphql.SchemaConfig{
		Query: types.RootQuery,
	}
	// define your schema
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Erro create schema, error: %v", err)
	}

	// define your graphql handler using graphql-go handler
	gqlHandler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// adapts and serve standard http handler
	http.Handle("/gqlhandler", gqlHandler)
	http.ListenAndServe(":8000", nil)
}
