package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

// Define Book type
var bookType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: graphql.String,
			},
			"publishedYear": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

// Mock data
var books = []map[string]interface{}{
	{"id": "1", "title": "1984", "author": "George Orwell", "publishedYear": 1949},
	{"id": "2", "title": "To Kill a Mockingbird", "author": "Harper Lee", "publishedYear": 1960},
	{"id": "3", "title": "The Catcher in the Rye", "author": "J. D. Salinger", "publishedYear": 1951},
	{"id": "4", "title": "The Great Gatsby", "author": "F. Scott Fitzgerald", "publishedYear": 1925},
}

// Define Query type
var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"books": &graphql.Field{
				Type: graphql.NewList(bookType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return books, nil
				},
			},
			"book": &graphql.Field{
				Type: bookType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.ID,
					}, 
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(string)
					if !ok {
						return nil, nil
					}
					for _, book := range books {
						if book["id"] == id {
							return book, nil
						}
					}
					return nil, nil
				},
			},
		},
	},
)

// Define Mutation type
var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"addBook": &graphql.Field{
				Type: bookType,
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"author": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"publishedYear": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					newBook := map[string]interface{}{
						"id":            fmt.Sprintf("%d", len(books)+1),
						"title":         p.Args["title"].(string),
						"author":        p.Args["author"].(string),
						"publishedYear": p.Args["publishedYear"].(int),
					}
					books = append(books, newBook)
					return newBook, nil
				},
			},
		},
	},
)

// Schema
var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)



func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		log.Printf("Errors: %+v", result.Errors)
	}
	return result
}

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var query struct {
			Query string `json:"query"`
		}

		err := json.NewDecoder(r.Body).Decode(&query)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		result := executeQuery(query.Query, schema)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	log.Println("Server is running on http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
