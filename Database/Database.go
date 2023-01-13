package Database

import "BookServer/Models"

// DataBase Stores the User and Book information
type DataBase struct {
	users map[string]Models.User
	books map[int]Models.Book
}

// NewDB() method creates a new DataBase instance
//func (db *DataBase) NewDB() *DataBase {
//
//}

// getUsers() method returns all the users
func (db *DataBase) getUsers() []Models.User {
	var allUsers []Models.User
	for _, user := range db.users {
		allUsers = append(allUsers, user)
	}
	return allUsers
}

// getUserByUserName() returns the user specified by param{username}
func (db *DataBase) getUserByUserName(username string) Models.User {
	var user Models.User
	user = db.users[username]
	return user
}

// getBooks() method returns all the books
func (db *DataBase) getBooks() []Models.Book {
	var allBooks []Models.Book
	for _, book := range db.books {
		allBooks = append(allBooks, book)
	}
	return allBooks
}

// getBookByBookId() returns the book specified by param{bookId}
func (db *DataBase) getBookByBookId(bookId int) Models.Book {
	var book Models.Book
	book = db.books[bookId]
	return book
}
