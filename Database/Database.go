// Package Database implements a basic database module for the Book-server.
// It provides a global 'db' object of type DataBase
// which is the central data storage for the Book-server.
// The db object can be created from backup files using NewDB function,
// and the pointer instance of the db object can be got by GetDB function.
package Database

import (
	"BookServer/Models"
	"BookServer/Utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/procyon-projects/chrono"
	"log"
	"math"
	"os"
	"sync"
	"time"
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
	// jobDelay is used for running a DataBase task after every certain amount of time.
	// Here jobDelay is used to back up DataBase data in backup folder after every jobDelay second.
	jobDelay time.Duration
}

// NewDB function creates a new DataBase instance with the backup data,
// Assigns into db instance and returns its address
// In addition it activates a Task Scheduler to perform
// data collection and store those data to back up folders.
func NewDB() *DataBase {
	db = DataBase{
		Users:      NewUserType(),
		Books:      NewBookType(),
		nextUserId: 1000,
		nextBookId: 1000,
		jobDelay:   time.Second * 20,
	}
	// Restoring data
	uJsonData, bJsonData := Utils.RestoreDataFromBackupFiles()
	var us []*Models.User
	var bs []*Models.Book

	err := json.Unmarshal(uJsonData, &us)
	err = json.Unmarshal(bJsonData, &bs)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// Assign restored data to db and update nextUserId, nextBookId
	for _, u := range us {
		db.nextUserId = int(math.Max(float64(u.Id), float64(db.nextUserId)+1))
		db.Users[u.Username] = u
	}
	for _, b := range bs {
		db.nextBookId = int(math.Max(float64(b.Id), float64(db.nextBookId)+1))
		db.Books[b.Id] = b
	}
	// Activate the Backup Scheduler
	// It will back up DataBase data after every
	// certain amount of time.
	db.DbBackupScheduler()

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

// DbBackupScheduler performs a task to store the data in the backup folder
// after every db.jobDelay. The default time delay is 20 seconds.
func (d *DataBase) DbBackupScheduler() {
	taskScheduler := chrono.NewDefaultTaskScheduler()
	_, err := taskScheduler.ScheduleWithFixedDelay(func(ctx context.Context) {
		// Decompose database data
		us, bs := d.decomposeDatabaseData()

		usersJson, err := json.Marshal(us)
		booksJson, err := json.Marshal(bs)
		if err != nil {
			log.Print(err.Error())
			return
		}
		err = Utils.StoreDataToBackupFiles(usersJson, booksJson)
		if err != nil {
			log.Print(err.Error())
			return
		}
		log.Print("Scheduled Backup Successful")
	}, db.jobDelay)

	if err != nil {
		log.Print("Task scheduler didn't successfully started")
		os.Exit(1)
	}
	log.Print("Task has been scheduled successfully.")
}

// decomposeDatabaseData creates []Models.User and []Models.Book and returns
func (d *DataBase) decomposeDatabaseData() ([]Models.User, []Models.Book) {
	// Decompose database data
	var us []Models.User
	var bs []Models.Book
	d.Lock()
	for _, u := range d.Users {
		us = append(us, *u)
	}
	for _, b := range d.Books {
		bs = append(bs, *b)
	}
	d.UnLock()

	return us, bs
}
