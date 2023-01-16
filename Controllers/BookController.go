package Controllers

import (
	"BookServer/Database"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"strconv"
)

// GetBookList returns all the book list
func GetBookList(w http.ResponseWriter, r *http.Request) {
	db := Database.GetDB()
	db.Lock()
	defer db.UnLock()

	books, err := db.GetBooks()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("content-type", "application/json")
	w.Write(books)
}

// GetBook returns a specific book information associated with id
func GetBook(w http.ResponseWriter, r *http.Request) {
	db := Database.GetDB()
	db.Lock()
	defer db.UnLock()

	bookId, err := strconv.Atoi(chi.URLParam(r, "id"))
	book, err := db.GetBookByBookId(bookId)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("content-type", "application/json")
	w.Write(book)
}

// CreateBook creates a book
// Returns the created book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	db := Database.GetDB()
	db.Lock()
	defer db.UnLock()

	body, err := io.ReadAll(r.Body)
	book, err := db.CreateBook(body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("content-type", "application/json")
	w.Write(book)
}

// DeleteBook deletes a book specified by the param{bookId}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// todo: .....
	w.Write([]byte(r.URL.Path))
}
