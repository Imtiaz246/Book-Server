package Models

import "time"

type Book struct {
	Id          int
	isbn        string
	name        string
	authors     []*User
	updatedAt   time.Time
	publishDate time.Time
	BookContent bookContent
}

type bookContent struct {
	chapters []chapter
}

type chapter struct {
	chapterTitle   string
	chapterContent string
}
