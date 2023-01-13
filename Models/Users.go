package Models

import "time"

type User struct {
	Id        int
	role      string
	username  string
	password  string
	createdAt time.Time
	updatedAt time.Time
}
