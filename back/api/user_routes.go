package api

import (
	"tbd/internal/rest/user"

	"github.com/go-chi/chi/v5"
	"github.com/uptrace/bun"
)

func UserRoutes(db *bun.DB)  chi.Router {
	r := chi.NewRouter()

	userHandler := user.UserHandler{}
	userHandler.DB = db

	r.Get("/", userHandler.ListUser)
	r.Post("/", userHandler.CreateUser)
	r.Get("/{id}", userHandler.GetUser)
	r.Put("/{id}", userHandler.UpdateUser)
	r.Delete("/{id}", userHandler.DeleteUser)
	return r
}