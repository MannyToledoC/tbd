package api

import (
	"tbd/internal/rest/user"

	"github.com/go-chi/chi/v5"
)

func UserRoutes()  chi.Router {
	r := chi.NewRouter()
	userHandler := user.UserHandler{}
	
	r.Get("/", userHandler.ListBooks)
	r.Post("/", userHandler.CreateBook)
	r.Get("/{id}", userHandler.GetBook)
	r.Put("/{id}", userHandler.UpdateBook)
	r.Delete("/{id}", userHandler.DeleteBook)
	return r
}