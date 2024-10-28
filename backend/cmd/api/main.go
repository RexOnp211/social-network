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

var DB *sql.DB

func main() {

	dbPath := "../../pkg/db/database.db"
	// Correct database path
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		_, err := os.Create(dbPath)
		if err != nil {
			fmt.Println("Error creating db file:", err)
			return
		}
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println("DB Open Error:", err)
		return
	}
	DB = db

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		fmt.Println("Driver Error:", err)
		return
	}

	migrationPath := "file://../../pkg/db/migrations/sqlite"

	// Correct migration path
	m, err := migrate.NewWithDatabaseInstance(
		migrationPath, "sqlite3", driver)
	if err != nil {
		fmt.Println("Migration Instance Error:", err)
		return
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println("Migration Up Error:", err)
		return
	}

	r := &api.Router{}
	r.AddRoute("GET", "/post/", http.HandlerFunc(handlers.HandlePosts))
	r.AddRoute("POST", "/post/", http.HandlerFunc(handlers.CreateComment))
	r.AddRoute("GET", "/", http.HandlerFunc(handlers.HomeHandler))
	r.AddRoute("GET", "/profile/", http.HandlerFunc(handlers.ProfileHandler))
	r.AddRoute("GET", "/group/", http.HandlerFunc(handlers.GroupHandler))
	r.AddRoute("POST", "/", http.HandlerFunc(handlers.CreatePostHandler))
	r.AddRoute("POST", "/register", http.HandlerFunc(handlers.RegisterUser))
	r.AddRoute("GET", "/image/", http.HandlerFunc(handlers.GetImageHandler))

	fmt.Println("Starting Go server")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Server Error:", err)
	}
}
