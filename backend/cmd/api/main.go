package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"social-network/internal/api"
	"social-network/internal/api/handlers"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Ensure the database file exists
	if _, err := os.Stat("../../pkg/db/database.db"); os.IsNotExist(err) {
		fmt.Println("Database file does not exist:", err)
		return
	}

	// Correct database path
	db, err := sql.Open("sqlite3", "../../pkg/db/database.db")
	if err != nil {
		fmt.Println("DB Open Error:", err)
		return
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		fmt.Println("Driver Error:", err)
		return
	}

	// Correct migration path
	m, err := migrate.NewWithDatabaseInstance(
		"file://../../pkg/db/migrations/sqlite", "sqlite3", driver)
	if err != nil {
		fmt.Println("Migration Instance Error:", err)
		return
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println("Migration Up Error:", err)
		return
	}

	r := &api.Router{}
	r.AddRoute("GET", "/posts", http.HandlerFunc(handlers.HandlePosts))
	r.AddRoute("GET", "/", http.HandlerFunc(handlers.HomeHandler))
	r.AddRoute("POST", "/", http.HandlerFunc(handlers.CreatePostHandler))

	fmt.Println("Starting Go server")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Server Error:", err)
	}
}
