// Package Database implements a basic database module for the Book-server.
// It provides a global 'db' object of type DataBase
// which is the central data storage for the Book-server.
// The db object can be created from backup files using NewDB function,
// and the pointer instance of the db object can be got by GetDB function.
package Database

import (
	"sync"
)

// DataBase Stores the User and Book information.
// As map values are inconsistent in the memory, pointers
// of each entity with corresponding key is stored in a map.
// Using pointer of UserType & BookType entity with respect to
// map kay is safe because server will run concurrently and multiple
// copy of DataBase object will lead to some dirty data.
type DataBase struct {
	Users                  UserType
	Books                  BookType
	nextUserId, nextBookId int
	// mu is used for synchronizing operations on the types.
	// As maps in golang are not concurrency proof,
	// So before doing [insert key, update key, get next id] operations,
	// we've to synchronize the operations.
	mu sync.Mutex
}

// NewDB function creates a new DataBase instance with the backup data,
// Assigns into db instance and returns its address
func NewDB() *DataBase {
	db = DataBase{
		Users:      NewUserType(),
		Books:      NewBookType(),
		nextUserId: 1001,
		nextBookId: 1001,
	}
	// todo: add recovery
	//uJsonData, bJsonData := Utils.RestoreDataFromBackupFiles()
	//
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

// GetNextUserId returns the next user id available to use
func (d *DataBase) GetNextUserId() int {
	return d.nextUserId
}

// GetNextBookId returns the next book id available to use
func (d *DataBase) GetNextBookId() int {
	return d.nextBookId
}

// Lock locks the database for synchronizing the operations
func (d *DataBase) Lock() {
	d.mu.Lock()
}

// UnLock unlocks the database for further operations
func (d *DataBase) UnLock() {
	d.mu.Unlock()
}
