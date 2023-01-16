package Database

// CreateUser creates a user and returns its json instance
func (d *DataBase) CreateUser(body []byte) ([]byte, error) {
	newUser, err := d.Users.Insert(body)
	return newUser, err
}

// GetUsers method returns all the users
func (d *DataBase) GetUsers() ([]byte, error) {
	allUsers, err := d.Users.Gets()
	return allUsers, err
}

// GetUserByUserName returns the user specified by param{username}
func (d *DataBase) GetUserByUserName(username string) ([]byte, error) {
	user, err := d.Users.Get(username)
	return user, err
}

// DeleteUserByUserName deletes a user record from the database.
// Returns error if any error occurs otherwise return nil
func (d *DataBase) DeleteUserByUserName(username string) error {
	_, err := d.Users.Get(username)
	if err != nil {
		return err
	}
	delete(d.Users, username)
	return nil
}

// CreateBook creates a book and returns its json object
func (d *DataBase) CreateBook(body []byte) ([]byte, error) {
	newBook, err := d.Books.Insert(body)
	return newBook, err
}

// GetBooks method returns all the books.
func (d *DataBase) GetBooks() ([]byte, error) {
	allBooks, err := d.Books.Gets()
	return allBooks, err
}

// GetBookByBookId returns book information specified by param{bookId}
func (d *DataBase) GetBookByBookId(bookId int) ([]byte, error) {
	book, err := d.Books.Get(bookId)
	return book, err
}

// DeleteBookByBookId returns book information specified by param{bookId}
func (d *DataBase) DeleteBookByBookId(bookId int) error {
	_, err := d.Books.Get(bookId)
	if err != nil {
		return err
	}
	delete(d.Books, bookId)
	return nil
}