package Database

import (
	"BookServer/Models"
	"BookServer/Utils"
	"errors"
	"time"
)

// UserType is a custom type which maps kay(username), value(*Model.User)
type UserType map[string]*Models.User

// BookType is a custom type which maps key(bookId), value(*Model.Book)
type BookType map[int]*Models.Book

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
	user, err := Models.NewUser(body)
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
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt

	(*u)[username] = user
	uJson, err := Utils.CreateSuccessJson(user)
	return uJson, err
}

// Gets returns all the user from the database in the json format.
// In addition, sends error if any occurs.
func (u *UserType) Gets() ([]byte, error) {
	var users []*Models.User
	for _, user := range *u {
		user.Password = ""
		users = append(users, user)
	}
	uJson, err := Utils.CreateSuccessJson(users)
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
	uJson, err := Utils.CreateSuccessJson(user)
	return uJson, err
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

// Insert inserts a book record into the database.
// In addition, it also checks for information validity.
// If the information is not valid it returns (nil, err) tuple,
// otherwise returns (json book object, nil) tuple.
func (b *BookType) Insert(body []byte) ([]byte, error) {
	book, err := Models.NewBook(body)
	if err != nil || !book.CheckValidity() {
		if err == nil {
			err = errors.New("book information is not valid")
		}
		return nil, err
	}
	book.Id = db.nextBookId
	db.nextBookId++
	book.PublishDate = time.Now()
	book.UpdatedAt = book.PublishDate

	(*b)[book.Id] = book
	bJson, err := Utils.CreateSuccessJson(book)
	return bJson, err
}

// Gets return all the book from the database without book-content
func (b *BookType) Gets() ([]byte, error) {
	var books []*Models.Book
	for _, book := range *b {
		book.BookContent = Models.BookContent{}
		books = append(books, book)
	}
	bJson, err := Utils.CreateSuccessJson(books)
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
	bJson, err := Utils.CreateSuccessJson(book)
	return bJson, err
}
