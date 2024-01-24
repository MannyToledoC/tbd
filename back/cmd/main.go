package main

import (
	"fmt"
	"net/http"
	"tbd/api"
	"tbd/internal/database/connections"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/extra/bundebug"
)

var db *bun.DB

func main() {
	fmt.Println("Hello, World")
	// ctx := context.Background()
	
	db = databaseConnection()
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	defer db.Close()
	// _, err := DB.NewCreateTable().Model((*models.User)(nil)).Exec(ctx)
	// if err != nil {
	// 	panic(err)
	// }
	routing()
}

func databaseConnection() *bun.DB{
	db := connections.Connect()
	return db
}

func routing() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("OK")) })
	r.Mount("/user", api.UserRoutes(db))
	http.ListenAndServe(":8080", r)
}