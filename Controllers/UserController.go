package Controllers

import (
	"encoding/json"
	"errors"
	"github.com/Imtiaz246/Book-Server/Database"
	"github.com/Imtiaz246/Book-Server/Models"
	"github.com/Imtiaz246/Book-Server/Utils"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

// GetUserList returns all the User list in a json format
func GetUserList(w http.ResponseWriter, _ *http.Request) {
	db := Database.GetDB()
	db.Lock()
	defer db.UnLock()

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

	user, err := db.GetUserByUserName(chi.URLParam(r, "username"))
	if err != nil {
		w.Write(user)
		return
	}
	w.Write(user)
}

// CreateUser creates a User. Returns the created User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := Database.GetDB()
	db.Lock()
	defer db.UnLock()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(Utils.CreateErrorJson(err))
		return
	}
	user, err := db.CreateUser(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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

	dUser := chi.URLParam(r, "username")
	// Checks if the request comes from admin otherwise returns bad request
	rUser := r.Context().Value("username").(string)
	if db.Users[rUser].Role != "admin" {
		w.WriteHeader(http.StatusForbidden)
		w.Write(Utils.CreateErrorJson(errors.New("don't have permission to delete a user")))
		return
	}
	// if rUser == dUser then it's admin
	if rUser == dUser {
		w.WriteHeader(http.StatusForbidden)
		w.Write(Utils.CreateErrorJson(errors.New("can't delete the admin")))
		return
	}
	err := db.DeleteUserByUserName(dUser)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(Utils.CreateErrorJson(err))
		return
	}
	msg, err := Utils.CreateSuccessJson("deleted successfully")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(Utils.CreateErrorJson(err))
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write(msg)
}

// GetToken generates a jwt token for a user.
// First it tries to authenticate the user credentials.
// If the credentials is valid then it generates a jwt token and returns it.
func GetToken(w http.ResponseWriter, r *http.Request) {
	db := Database.GetDB()
	db.Lock()
	defer db.UnLock()
	// User Credentials struct
	var uc struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	body, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(body, &uc)
	if err != nil {
		w.Write(Utils.CreateErrorJson(err))
		return
	}

	// Check if the user credentials is valid or not
	err = db.Authenticate(uc.Username, uc.Password)
	if err != nil {
		w.Write(Utils.CreateErrorJson(err))
		return
	}

	// Generate token and respond it
	tokenStr, err := Utils.GenerateJwtToken(uc.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(Utils.CreateErrorJson(err))
		return
	}
	msg, err := Utils.CreateSuccessJson(tokenStr)
	if err != nil {
		w.Write(Utils.CreateErrorJson(err))
		return
	}
	w.Write(msg)
}

// GetBooksOfUser returns the book list of a user defined by param{username}
func GetBooksOfUser(w http.ResponseWriter, r *http.Request) {
	db := Database.GetDB()
	db.Lock()
	defer db.UnLock()

	books := db.Users[chi.URLParam(r, "username")].BookOwns
	for _, b := range books {
		b.Authors = []*Models.User{}
	}
	msg, err := Utils.CreateSuccessJson(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(Utils.CreateErrorJson(err))
		return
	}
	w.Write(msg)
}
