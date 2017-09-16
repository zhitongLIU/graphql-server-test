package main

import (
	"net/http"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"latestPost": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Hello World!", nil
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})


func main() {
	// http handler with schema
	h := handler.New(&handler.Config {
		Schema: &Schema,
		Pretty: true,
	})

	// define endpoint with handler
	http.Handle("/graphql", h)

	http.ListenAndServe(":8080", nil)
}
