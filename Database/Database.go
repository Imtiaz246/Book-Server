// Package Database implements a basic database module for the Book-server.
// It provides a global 'db' object of type DataBase
// which is the central data storage for the Book-server.
// The db object can be created from backup files using NewDB function,
// and the pointer instance of the db object can be got by GetDB function.
package Database

import "BookServer/Models"

// DataBase Stores the User and Book information
// As map values are inconsistent in the memory,
// Pointers of each entity with corresponding key is stored in a map
type DataBase struct {
	Users map[string]*Models.User
	Books map[int]*Models.Book
}

// NewDB function creates a new DataBase instance with the backup data,
// Assigns into db instance and returns its address
func NewDB() *DataBase {
	db = DataBase{
		Users: make(map[string]*Models.User),
		Books: make(map[int]*Models.Book),
	}
	return &db
}

// GetDB returns the address of db instance
// which is the Global DataBase object
// or DataBase storage.
func GetDB() *DataBase {
	return &db
}

// db is the Central/Global DataBase instance
var db DataBase
