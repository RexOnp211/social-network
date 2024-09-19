package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"social-network/internal/api"
	"social-network/internal/api/handlers"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, err := sql.Open("sqlite3", "/backend/pkg/db/database.db")
	if err != nil {
		fmt.Println(err)
	}
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		fmt.Println(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"sqlite3", driver)
	m.Up() // or m.Steps(2) if you want to explicitly set the number of migrations to run
	if err != nil {
		fmt.Println(err)
	}

	r := &api.Router{}

	r.AddRoute("GET", "/posts", http.HandlerFunc(handlers.HandlePosts))
	r.AddRoute("GET", "/", http.HandlerFunc(handlers.HomeHandler))

	fmt.Println("starting go server")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
	}
}
