package Controllers

import (
	"BookServer/Models"
	"fmt"
	"io"
	"net/http"
)

// GetBookList returns all the book list
func GetBookList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}

// GetBook returns a specific book information associated with id
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}

// CreateBook creates a book
// Returns the created book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	jsonObj, _ := io.ReadAll(r.Body)
	newBook, err := Models.NewBook(jsonObj)
	if err != nil {
		fmt.Println("kire baler error", err)
		w.Write([]byte("baer erroe: "))
	}
	fmt.Println("new book si : ", *newBook)
	jsonObj, _ = newBook.GenerateJSON()
	w.Write([]byte(jsonObj))
}

// DeleteBook deletes a book specified by the param{bookId}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}
