package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

// getBookList() returns all the book list
func getBookList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}

// getBook() returns a specific book information associated with id
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}

// createBook() creates a book
func createBook(w http.ResponseWriter, r *http.Request) {
	c, _ := io.ReadAll(r.Body)
	fmt.Println(string(c))
	w.Write([]byte(c))
}

// deleteBook() deletes a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}

func main() {
	// New chi Router
	r := chi.NewRouter()

	// Middleware lists
	r.Use(middleware.Logger)

	// APIS
	r.Get("/api/v1/books", getBookList)
	r.Get("/api/v1/books/{id}", getBook)
	r.Post("/api/v1/books", createBook)
	r.Delete("/api/v1/books/{id}", deleteBook)

	// Start the server and Listen on port :3000
	http.ListenAndServe(":3000", r)
}
