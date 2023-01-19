package Models

import (
	"encoding/json"
	"time"
)

type User struct {
	Id           int       `json:"id"`
	Role         string    `json:"role"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	Password     string    `json:"password,omitempty"`
	Organization string    `json:"organization"`
	BookOwns     []*Book   `json:"book-owns,omitempty"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

// NewUser creates a user instance, from the http request body.
// Returns the pointer of the new user instance
func NewUser(body []byte) (*User, error) {
	var newUser User
	err := json.Unmarshal(body, &newUser)
	return &newUser, err
}

// CheckValidity checks if user information is valid or not. If not valid,
// it returns False, otherwise returns True.
func (u *User) CheckValidity() bool {
	nl, pl := len(u.Username), len(u.Password)
	if nl == 0 || pl == 0 {
		return false
	}
	return true
}

// GenerateJSON generates JSON object from Go object
// Returns the (json object, error) tuple
func (u *User) GenerateJSON() ([]byte, error) {
	jsonObj, err := json.Marshal(u)
	return jsonObj, err
}
