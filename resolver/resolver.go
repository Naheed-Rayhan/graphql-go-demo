package resolver

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

// Mock data
var books = []map[string]interface{}{
	{"id": "1", "title": "1984", "author": "George Orwell", "publishedYear": 1949},
	{"id": "2", "title": "To Kill a Mockingbird", "author": "Harper Lee", "publishedYear": 1960},
	{"id": "3", "title": "The Catcher in the Rye", "author": "J. D. Salinger", "publishedYear": 1951},
	{"id": "4", "title": "The Great Gatsby", "author": "F. Scott Fitzgerald", "publishedYear": 1925},
}

func GetBooks(p graphql.ResolveParams) (interface{}, error) {
	return books, nil
}

func GetBookByID(p graphql.ResolveParams) (interface{}, error) {
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
}

func AddBook(p graphql.ResolveParams) (interface{}, error) {
	newBook := map[string]interface{}{
		"id":            fmt.Sprintf("%d", len(books)+1),
		"title":         p.Args["title"].(string),
		"author":        p.Args["author"].(string),
		"publishedYear": p.Args["publishedYear"].(int),
	}
	books = append(books, newBook)
	return newBook, nil
}