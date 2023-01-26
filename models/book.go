package models

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
	Authors     []*User     `json:"authors,omitempty"`
	UpdatedAt   time.Time   `json:"updated-at,omitempty"`
	PublishDate time.Time   `json:"publish-date,omitempty"`
	BookContent BookContent `json:"book-content,omitempty"`
}

type BookContent struct {
	OverView string    `json:"over-view,omitempty"`
	Chapters []chapter `json:"chapters,omitempty"`
}

type chapter struct {
	ChapterTitle   string `json:"chapter-title"`
	ChapterContent string `json:"chapter-content"`
}

// NewBook creates a new book instance, from the json []byte slice.
// Returns the address of the book instance
func NewBook(body []byte) (*Book, error) {
	var newBook Book
	err := json.Unmarshal(body, &newBook)
	return &newBook, err
}

// CheckValidity checks if book information is valid or not. If not valid,
// it returns False, otherwise returns True.
func (b *Book) CheckValidity() bool {
	nl, al, cl := len(b.BookName), len(b.Authors), len(b.BookContent.Chapters)
	if nl == 0 || al == 0 || cl == 0 {
		return false
	}
	return true
}

// GenerateJSON creates JSON object from Book object.
// Returns the JSON []byte object, error tuple
func (b *Book) GenerateJSON() ([]byte, error) {
	jsonObj, err := json.Marshal(b)
	return jsonObj, err
}
