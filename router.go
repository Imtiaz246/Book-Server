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
	r.Route("/api/v1/books", bookRoutes)
	// User APIS
	r.Route("/api/v1/users", userRoutes)

	return r
}

func bookRoutes(r chi.Router) {
	// Protected Routes, Authentication required
	r.Group(func(r chi.Router) {
		// Todo: apply jwt authentication
		r.Post("/", Controllers.CreateBook)
		r.Delete("/{id}", Controllers.DeleteBook)
	})
	// Public Routes
	r.Get("/", Controllers.GetBookList)
	r.Get("/{id}", Controllers.GetBook)
}

func userRoutes(r chi.Router) {
	// Protected Routes, Authentication required
	r.Group(func(r chi.Router) {
		// Todo: apply jwt authentication
		r.Post("/", Controllers.CreateUser)
		r.Delete("/{id}", Controllers.DeleteUser)
	})
	// Public Routes
	r.Get("/", Controllers.GetUserList)
	r.Get("/{id}", Controllers.GetUser)
}
