package database

import (
	"encoding/json"
	"errors"
	"github.com/Imtiaz246/Book-Server/models"
	"github.com/Imtiaz246/Book-Server/utils"
	"time"
)

// UserType is a custom type which maps kay(username), value(*Model.User)
type UserType map[string]*models.User

// BookType is a custom type which maps key(bookId), value(*Model.Book)
type BookType map[int]*models.Book

// NewUserType creates a new UserType and returns the instance
func NewUserType() UserType {
	newUserType := make(UserType)
	return newUserType
}

// NewBookType creates a new BookType and returns the instance
func NewBookType() BookType {
	newBookType := make(BookType)
	return newBookType
}

// Insert tries to add a user to the database. If a username already exists,
// it returns a (nil, error) tuple otherwise returns the (json user object, nil) tuple
func (u *UserType) Insert(body []byte) ([]byte, error) {
	user, err := models.NewUser(body)
	if err != nil || !user.CheckValidity() {
		if err == nil {
			err = errors.New("user information is not valid")
		}
		return nil, err
	}
	username := user.Username
	if _, found := (*u)[username]; found == true {
		err = errors.New("duplicate username")
		return nil, err
	}
	user.Id = db.nextUserId
	db.nextUserId++
	user.Role = "user"
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt

	(*u)[username] = user
	uJson, err := utils.CreateSuccessJson(user)
	return uJson, err
}

// Gets returns all the user from the database in the json format.
// In addition, sends error if any occurs.
func (u *UserType) Gets() ([]byte, error) {
	var users []*models.User
	for _, user := range *u {
		cUser := *user
		cUser.Password = ""
		users = append(users, &cUser)
	}
	uJson, err := utils.CreateSuccessJson(users)
	return uJson, err
}

// Get returns a specific user defined by the param{username}.
// If a record is found, it returns the (json user object, nil) tuple,
// otherwise returns (nil, err) tuple.
func (u *UserType) Get(username string) ([]byte, error) {
	user, found := (*u)[username]
	if !found {
		err := errors.New("username doesn't exists")
		return nil, err
	}
	uJson, err := utils.CreateSuccessJson(user)
	return uJson, err
}

// DeleteUser deletes a user from the DataBase
func (u *UserType) DeleteUser(username string) error {
	_, found := (*u)[username]
	if !found {
		return errors.New("username not found")
	}
	delete(*u, username)
	return nil
}

// UpdateUser updates a user from the DataBase
func (u *UserType) UpdateUser(username string, body []byte) error {
	usr, found := (*u)[username]
	if !found {
		return errors.New("username not found")
	}
	var tu models.User
	var err error
	err = json.Unmarshal(body, &tu)
	if err != nil {
		return err
	}
	tu.Id = usr.Id
	if !tu.CheckValidity() {
		return errors.New("user information is not valid")
	}
	(*u)[username] = &tu
	return nil
}

// CreateAdmin creates an admin for the server.
// Admin has permission to delete user and also books.
// Admin is created with username "imtiaz" and password "1234"
func (u *UserType) CreateAdmin() error {
	_, found := (*u)["imtiaz"]
	if found {
		return nil
	}
	admin := models.User{
		Id:           100,
		Username:     "imtiaz",
		Password:     "1234",
		Organization: "Appscode Ltd",
		Email:        "imtiazuddincho246@gmail.com",
		Role:         "admin",
	}
	(*u)[admin.Username] = &admin
	_, found = (*u)[admin.Username]
	if !found {
		return errors.New("can't create admin")
	}
	return nil
}

// CheckCredentials checks if a user credentials are exists in the database of not.
// If not valid or doesn't exist it returns an error.
func (u *UserType) CheckCredentials(username, password string) error {
	user, found := (*u)[username]
	if !found {
		err := errors.New("username doesn't exists")
		return err
	}
	if user.Password != password {
		err := errors.New("password doesn't match")
		return err
	}
	return nil
}

// UsersExistence checks if a list of user is exists in DataBase or not.
// if not exists returns an error.
func (u *UserType) UsersExistence(users []*models.User) error {
	for _, usr := range users {
		_, found := (*u)[usr.Username]
		if !found {
			return errors.New("username " + usr.Username + " doesn't exists")
		}
	}
	return nil
}

// Insert inserts a book record into the database.
// In addition, it also checks for information validity & authors validation &
// update the BookOwns property of User field of DataBase.
// If the information is not valid it returns (nil, err) tuple,
// otherwise returns (json book object, nil) tuple.
func (b *BookType) Insert(body []byte) ([]byte, error) {
	book, err := models.NewBook(body)
	if err != nil || !book.CheckValidity() {
		if err == nil {
			err = errors.New("book information is not valid")
		}
		return nil, err
	}
	// checks authors existence in the DataBase
	// if any username is not exists in the db it returns an error
	err = db.AuthenticateUsersExistence(book.Authors)
	if err != nil {
		return nil, err
	}
	// Connect authors with the book created
	for _, author := range book.Authors {
		db.Users[author.Username].BookOwns = append(db.Users[author.Username].BookOwns, book)
	}

	book.Id = db.nextBookId
	db.nextBookId++
	book.PublishDate = time.Now()
	book.UpdatedAt = book.PublishDate

	(*b)[book.Id] = book
	bJson, err := utils.CreateSuccessJson(book)
	return bJson, err
}

// Gets return all the book from the database without book-content
func (b *BookType) Gets() ([]byte, error) {
	var books []*models.Book
	for _, book := range *b {
		book.BookContent = models.BookContent{}
		books = append(books, book)
	}
	bJson, err := utils.CreateSuccessJson(books)
	return bJson, err
}

// Get returns a specific book defined by the param{bookId}.
// If a record is found, it returns the (json book object, nil) tuple,
// otherwise returns (nil, error) tuple.
func (b *BookType) Get(bookId int) ([]byte, error) {
	book, found := (*b)[bookId]
	if !found {
		err := errors.New("book doesn't exists")
		return nil, err
	}
	bJson, err := utils.CreateSuccessJson(book)
	return bJson, err
}

// DeleteBook deletes a book from the DataBase
func (b *BookType) DeleteBook(bookId int, requestedUser string) error {
	book, found := (*b)[bookId]
	if !found {
		return errors.New("book doesn't exists")
	}
	for _, a := range book.Authors {
		if a.Username == requestedUser {
			goto DELETE
		}
	}
	return errors.New("user don't have permission")
DELETE:
	delete(*b, bookId)
	return nil
}

// UpdateBook updates a book from the DataBase
func (b *BookType) UpdateBook(bookId int, requestedUser string, body []byte) error {
	book, found := (*b)[bookId]
	if !found {
		return errors.New("book doesn't exists")
	}
	for _, a := range book.Authors {
		if a.Username == requestedUser {
			goto UPDATE
		}
	}
	return errors.New("user don't have permission")
UPDATE:
	var tb models.Book
	err := json.Unmarshal(body, &tb)
	if err != nil {
		return err
	}
	if !tb.CheckValidity() {
		return errors.New("book information is not valid")
	}
	nid := book.Id
	tb.Id = nid

	(*b)[bookId] = &tb
	return nil
}
