package Controllers

import (
	"BookServer/Database"
	"BookServer/Utils"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"strconv"
)

// GetBookList returns all the book list
func GetBookList(w http.ResponseWriter, _ *http.Request) {
	db := Database.GetDB()
	db.Lock()
	defer db.UnLock()

	bookList, err := db.GetBooks()
	if err != nil {
		w.Write(Utils.CreateErrorJson(err))
		return
	}
	w.Write(bookList)
}

// GetBook returns a specific book information associated with id
func GetBook(w http.ResponseWriter, r *http.Request) {
	db := Database.GetDB()
	db.Lock()
	defer db.UnLock()

	bookId, err := strconv.Atoi(chi.URLParam(r, "id"))
	book, err := db.GetBookByBookId(bookId)
	if err != nil {
		w.Write(Utils.CreateErrorJson(err))
		return
	}
	w.Write(book)
}

// CreateBook creates a book. Returns the created book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	db := Database.GetDB()
	db.Lock()
	defer db.UnLock()

	body, err := io.ReadAll(r.Body)
	nBook, err := db.CreateBook(body)
	if err != nil {
		w.Write(Utils.CreateErrorJson(err))
		return
	}
	w.Write(nBook)
}

// DeleteBook deletes a book specified by the param{bookId}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	db := Database.GetDB()
	db.Lock()
	defer db.UnLock()
	// checks if the request comes from admin or author
	requestedUser := r.Context().Value("username").(string)

	bookId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(Utils.CreateErrorJson(err))
		return
	}
	err = db.DeleteBookByBookId(bookId, requestedUser)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write(Utils.CreateErrorJson(err))
		return
	}
	msg, err := Utils.CreateSuccessJson("deleted successfully")
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write(Utils.CreateErrorJson(err))
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write(msg)
}
