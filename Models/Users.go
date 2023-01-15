package Models

import (
	"encoding/json"
	"time"
)

type User struct {
	Id           int       `json:"id"`
	Role         string    `json:"role"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Organization string    `json:"organization"`
	BookOwns     []*Book   `json:"book-owns"`
	CreatedAt    time.Time `json:"created-at"`
	UpdatedAt    time.Time `json:"updated-at"`
}

// NewUser creates a user instance, from the json []byte slice.
// Returns the pointer of the user instance
func NewUser(jsonObj []byte) (*User, error) {
	var newUser User
	err := json.Unmarshal(jsonObj, &newUser)
	return &newUser, err
}

// GenerateJSON generates JSON object from Go object
// Returns the json object, error tuple
func (u *User) GenerateJSON() ([]byte, error) {
	jsonObj, err := json.Marshal(u)
	return jsonObj, err
}
