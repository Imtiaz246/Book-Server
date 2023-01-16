package Controllers

import (
	"BookServer/Database"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

// GetUserList returns all the User list in a json format
func GetUserList(w http.ResponseWriter, r *http.Request) {
	db := Database.GetDB()
	db.Lock()
	defer db.UnLock()

	userList, err := db.GetUsers()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("content-type", "application/json")
	w.Write(userList)
}

// GetUser returns a specific User information associated with id
func GetUser(w http.ResponseWriter, r *http.Request) {
	db := Database.GetDB()
	db.Lock()
	defer db.UnLock()

	user, err := db.GetUserByUserName(chi.URLParam(r, "username"))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("content-type", "application/json")
	w.Write(user)
}

// CreateUser creates a User
// Returns the created User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := Database.GetDB()
	db.Lock()
	defer db.UnLock()

	body, err := io.ReadAll(r.Body)
	user, err := db.CreateUser(body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("content-type", "application/json")
	w.Write(user)
}

// GetBooksOfUser returns the books list of a specific user defined by param{username}
func GetBooksOfUser(w http.ResponseWriter, r *http.Request) {
	// todo: ...
}

// DeleteUser deletes a User specified by the param{UserId}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := Database.GetDB()
	db.Lock()
	defer db.UnLock()

	err := db.DeleteUserByUserName(chi.URLParam(r, "username"))
	if err != nil {
		w.WriteHeader(204)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(202)
	w.Write([]byte("deleted successfully"))
}
