package Controllers

import (
	"BookServer/Database"
	"BookServer/Utils"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

// GetUserList returns all the User list in a json format
func GetUserList(w http.ResponseWriter, r *http.Request) {
	db := Database.GetDB()
	db.Lock()
	defer db.UnLock()
	w.Header().Add("content-type", "application/json")

	userList, err := db.GetUsers()
	if err != nil {
		w.Write(userList)
		return
	}
	w.Write(userList)
}

// GetUser returns a specific User information associated with id
func GetUser(w http.ResponseWriter, r *http.Request) {
	db := Database.GetDB()
	db.Lock()
	defer db.UnLock()
	w.Header().Add("content-type", "application/json")

	user, err := db.GetUserByUserName(chi.URLParam(r, "username"))
	if err != nil {
		w.Write(user)
		return
	}
	w.Write(user)
}

// CreateUser creates a User
// Returns the created User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := Database.GetDB()
	db.Lock()
	defer db.UnLock()
	w.Header().Add("content-type", "application/json")

	body, err := io.ReadAll(r.Body)
	user, err := db.CreateUser(body)
	if err != nil {
		w.Write(Utils.CreateErrorJson(err))
		return
	}
	w.Write(user)
}

// DeleteUser deletes a User specified by the param{UserId}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := Database.GetDB()
	db.Lock()
	defer db.UnLock()
	w.Header().Add("content-type", "application/json")

	err := db.DeleteUserByUserName(chi.URLParam(r, "username"))
	msg, err := Utils.CreateSuccessJson([]byte("deleted successfully"))
	if err != nil {
		w.WriteHeader(204)
		w.Write(Utils.CreateErrorJson(err))
		return
	}
	w.WriteHeader(202)
	w.Write(msg)
}
