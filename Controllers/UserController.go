package Controllers

import (
	"BookServer/Database"
	"BookServer/Utils"
	"encoding/json"
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
	msg, err := Utils.CreateSuccessJson(tokenStr)
	if err != nil {
		w.Write(Utils.CreateErrorJson(err))
		return
	}
	w.Write(msg)
}
