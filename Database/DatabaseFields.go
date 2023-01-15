package Database

import (
	"BookServer/Models"
	"errors"
)

type UserType map[string]*Models.User
type BookType map[int]*Models.Book

func (u *UserType) insert(user *Models.User) (*Models.User, error) {
	username := user.Username
	(*u)[username] = user
	return (*u)[username], nil
}

func (u *UserType) gets() []*Models.User {
	var users []*Models.User
	for _, user := range *u {
		users = append(users, user)
	}
	return users
}

func (u *UserType) get(username string) (*Models.User, error) {
	user, found := (*u)[username]
	if !found {
		err := errors.New("username doesn't exists")
		return nil, err
	}
	return user, nil
}
