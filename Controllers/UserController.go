package Controllers

import (
	"BookServer/Database"
	"BookServer/Models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetUserList returns all the User list
func GetUserList(w http.ResponseWriter, r *http.Request) {
	user := Models.User{
		Id:       1,
		Role:     "Admin",
		Username: "Imtiaz",
		Password: "1234",
	}
	user1 := Models.User{
		Id:       2,
		Role:     "Admin",
		Username: "ImtiazUddin",
		Password: "1234",
	}

	db := Database.NewDB()
	db.Users["Imtiaz"] = &user
	db.Users["uddin"] = &user1

	var users []Models.User
	for _, value := range db.Users {
		users = append(users, *value)
	}

	w.WriteHeader(200)
	res, _ := json.Marshal(users)
	w.Write(res)
}

// GetUser returns a specific User information associated with id
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}

// CreateUser creates a User
// Returns the created User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	c, _ := io.ReadAll(r.Body)
	fmt.Println(string(c))
	w.Write([]byte(c))
}

// DeleteUser deletes a User specified by the param{UserId}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}
