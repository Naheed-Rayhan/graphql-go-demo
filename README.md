
# Go GraphQL API for Books

This is a simple GraphQL API written in Go that allows you to query and mutate (add) a list of books. The server exposes a `/graphql` endpoint, where you can interact with the data using GraphQL queries and mutations.

## Features

- **Query** the list of books or a single book by its `id`.
- **Mutate** the list of books by adding a new book.
- Written in **Go** using the `graphql-go` package.

## Getting Started

Follow these steps to set up and run the project locally.

### Prerequisites

Make sure you have the following installed:

- **Go** (1.16 or later) - Download and install Go from [here](https://golang.org/dl/).
- **Git** - Download and install Git from [here](https://git-scm.com/).

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/go-graphql-books-api.git
   cd go-graphql-books-api
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

### Running the Server

To start the GraphQL API server, run the following command:

```bash
go run main.go
```

The server will start and listen on `http://localhost:8080/graphql`.

### Querying the API

You can send **POST requests** to `http://localhost:8080/graphql` with a GraphQL query in the request body. Below are some example queries and mutations.

#### Query: Get All Books

```graphql
{
  books {
    id
    title
    author
    publishedYear
  }
}
```

**Response:**

```json
{
  "data": {
    "books": [
      {"id": "1", "title": "1984", "author": "George Orwell", "publishedYear": 1949},
      {"id": "2", "title": "To Kill a Mockingbird", "author": "Harper Lee", "publishedYear": 1960},
      {"id": "3", "title": "The Catcher in the Rye", "author": "J. D. Salinger", "publishedYear": 1951},
      {"id": "4", "title": "The Great Gatsby", "author": "F. Scott Fitzgerald", "publishedYear": 1925}
    ]
  }
}
```

#### Query: Get a Single Book by ID

```graphql
{
  book(id: "2") {
    title
    author
    publishedYear
  }
}
```

**Response:**

```json
{
  "data": {
    "book": {
      "title": "To Kill a Mockingbird",
      "author": "Harper Lee",
      "publishedYear": 1960
    }
  }
}
```

#### Mutation: Add a New Book

```graphql
mutation {
  addBook(title: "Brave New World", author: "Aldous Huxley", publishedYear: 1932) {
    id
    title
    author
    publishedYear
  }
}
```

**Response:**

```json
{
  "data": {
    "addBook": {
      "id": "5",
      "title": "Brave New World",
      "author": "Aldous Huxley",
      "publishedYear": 1932
    }
  }
}
```

### Example Request Using `curl`

You can also interact with the API using `curl`. Here's an example:

```bash
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "{ books { id title author publishedYear } }"}'
```

### Testing the API

You can test the API using tools like:

- **GraphQL Playground**: Open [http://localhost:8080/graphql](http://localhost:8080/graphql) and interact with the API via a GraphQL Playground interface.
- **Postman**: Create a POST request to `http://localhost:8080/graphql` with the GraphQL query as the body.

## Code Structure

- **`main.go`**: The main entry point for the application, where the server is initialized and GraphQL schema is defined.
- **`books`**: A mock list of books used for querying and mutation.
- **`graphql-go`**: The package used for defining GraphQL types, queries, mutations, and executing the queries.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Feel free to fork this repository and create pull requests. Contributions are welcome!

---

For more information on GraphQL, you can visit the [official GraphQL website](https://graphql.org/).

