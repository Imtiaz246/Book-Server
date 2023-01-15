package Database

import "BookServer/Models"

// GetUsers method returns all the users
func (db *DataBase) GetUsers() []Models.User {
	var allUsers []Models.User
	for _, user := range db.Users {
		allUsers = append(allUsers, *user)
	}
	return allUsers
}

// GetUserByUserName returns the user specified by param{username}
func (db *DataBase) GetUserByUserName(username string) Models.User {
	var user Models.User
	user = *db.Users[username]
	return user
}

// DeleteUserByUserName deletes a user record from the database.
// Returns True if operation is successful,
// returns False otherwise
func (db *DataBase) DeleteUserByUserName(username string) (bool, string) {
	if db.Users[username] == nil {
		return false, "Record doesn't exists in the database!!!"
	}
	db.Users[username] = nil
	return true, "Record deleted successfully!!!"
}

// GetBooks method returns all the books.
func (db *DataBase) GetBooks() []Models.Book {
	var allBooks []Models.Book
	for _, book := range db.Books {
		allBooks = append(allBooks, *book)
	}
	return allBooks
}

// GetBookByBookId returns book information specified by param{bookId}
func (db *DataBase) GetBookByBookId(bookId int) Models.Book {
	var book Models.Book
	book = *db.Books[bookId]
	return book
}
