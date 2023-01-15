package Models

import (
	"encoding/json"
	"time"
)

type Book struct {
	Id          int         `json:"id"`
	Price       int         `json:"price"`
	SoldCopies  int         `json:"sold-copies"`
	Isbn        string      `json:"isbn"`
	BookName    string      `json:"book-name"`
	Authors     []*User     `json:"authors"`
	UpdatedAt   time.Time   `json:"updated-at"`
	PublishDate time.Time   `json:"publish-date"`
	BookContent bookContent `json:"book-content"`
}

type bookContent struct {
	OverView string    `json:"over-view"`
	Chapters []chapter `json:"chapters"`
}

type chapter struct {
	ChapterTitle   string `json:"chapter-title"`
	ChapterContent string `json:"chapter-content"`
}

// NewBook function creates a new book instance, from the json []byte slice.
// Returns the address of the book instance
func NewBook(jsonObj []byte) (*Book, error) {
	var newBook Book
	err := json.Unmarshal(jsonObj, &newBook)
	if err != nil {
		return &newBook, err
	}
	// TODO: generate id and other things that has to be auto generated

	return &newBook, err
}

// GenerateJSON creates JSON object from Book object.
// Returns the JSON []byte object, error tuple
func (b *Book) GenerateJSON() ([]byte, error) {
	jsonObj, err := json.Marshal(b)
	return jsonObj, err
}
