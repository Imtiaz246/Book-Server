package main

import (
	"BookServer/Controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Router() http.Handler {
	// New chi Router instance
	r := chi.NewRouter()

	// Middleware lists
	r.Use(middleware.Logger)

	// Book APIS
	r.Route("/api/v1/books", BookRoutes)
	// User APIS
	r.Route("/api/v1/users", UserRoutes)

	return r
}

func BookRoutes(r chi.Router) {
	// Protected Routes, Authentication required
	r.Group(func(r chi.Router) {
		// TODO: apply jwt authentication
		r.Post("/", Controllers.CreateBook)
		r.Delete("/{id}", Controllers.DeleteBook)
	})
	// Public Routes
	r.Get("/", Controllers.GetBookList)
	r.Get("/{id}", Controllers.GetBook)
}

func UserRoutes(r chi.Router) {
	// Protected Routes, Authentication required
	r.Group(func(r chi.Router) {
		// TODO: apply jwt authentication
		r.Post("/", Controllers.CreateUser)
		r.Delete("/{username}", Controllers.DeleteUser)
	})
	// Public Routes
	r.Get("/", Controllers.GetUserList)
	r.Get("/{username}", Controllers.GetUser)
	r.Get("/{username}/books", Controllers.GetBooksOfUser)
}
