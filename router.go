package main

import (
	"BookServer/Controllers"
	"BookServer/Middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Router() http.Handler {
	// New chi Router instance
	r := chi.NewRouter()

	// Middleware lists
	r.Use(middleware.Logger)

	// Add common Header for all routes
	r.Use(Middleware.AddHeaders)

	// Book APIS
	r.Route("/api/v1/books", BookRoutes)
	// User APIS
	r.Route("/api/v1/users", UserRoutes)

	return r
}

func BookRoutes(r chi.Router) {
	// Protected Routes, Authentication required
	r.Group(func(r chi.Router) {
		// Custom Auth middleware
		//r.Use(Middleware.BasicAuth)
		r.Use(Middleware.JwtAuth)

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
		// Custom Auth middleware
		//r.Use(Middleware.BasicAuth)
		r.Use(Middleware.JwtAuth)
		
		r.Delete("/{username}", Controllers.DeleteUser)
	})
	// Public Routes
	r.Post("/", Controllers.CreateUser)
	r.Get("/", Controllers.GetUserList)
	r.Get("/{username}", Controllers.GetUser)
	r.Get("/{username}/books", Controllers.GetBooksOfUser)
	r.Post("/get-token", Controllers.GetToken)
}
