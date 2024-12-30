package main

import (
	"encoding/json"
	"graphql-go-demo/schema"
	"log"
	"net/http"
)






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

		result := schema.ExecuteQuery(query.Query)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	log.Println("Server is running on http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
