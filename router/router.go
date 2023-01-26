package router

import (
	"github.com/Imtiaz246/Book-Server/controllers"
	"github.com/Imtiaz246/Book-Server/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Router() http.Handler {
	// New chi router instance
	r := chi.NewRouter()

	// middlewares lists
	r.Use(middleware.Logger)
	// Add common Header for all routes
	r.Use(middlewares.AddHeaders)

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
		//r.Use(middlewares.BasicAuth)
		r.Use(middlewares.JwtAuth)

		r.Post("/", controllers.CreateBook)
		r.Delete("/{id}", controllers.DeleteBook)
		r.Put("/{id}", controllers.UpdateBook)
	})
	// Public Routes
	r.Get("/", controllers.GetBookList)
	r.Get("/{id}", controllers.GetBook)
}

func UserRoutes(r chi.Router) {
	// Protected Routes, Authentication required
	r.Group(func(r chi.Router) {
		// Custom Auth middleware
		//r.Use(middlewares.BasicAuth)
		r.Use(middlewares.JwtAuth)

		r.Delete("/{username}", controllers.DeleteUser)
		r.Put("/{username}", controllers.UpdateUser)
	})
	// Public Routes
	r.Post("/", controllers.CreateUser)
	r.Get("/", controllers.GetUserList)
	r.Get("/{username}", controllers.GetUser)
	r.Get("/{username}/books", controllers.GetBooksOfUser)
	r.Post("/get-token", controllers.GetToken)
}
