package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"social-network/internal/api"
	"social-network/internal/api/handlers"
	db "social-network/pkg/db/sqlite"

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
	DB, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println("DB Open Error:", err)
		return
	}

	// Enable foreign key constraints
	_, err = DB.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		fmt.Println("Failed to enable foreign keys:", err)
		return
	}
	row := DB.QueryRow("PRAGMA foreign_keys;")
	var foreignKeyStatus int
	err = row.Scan(&foreignKeyStatus)
	if err != nil {
		fmt.Println("Error checking foreign key status:", err)
	} else {
		fmt.Printf("Foreign Key Status: %d\n", foreignKeyStatus)
	}

	db.SetDB(DB)

	driver, err := sqlite3.WithInstance(DB, &sqlite3.Config{})
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

	// if err := m.Down(); err != nil && err != migrate.ErrNoChange {
	// 	fmt.Println("Migration Down Error:", err)
	// 	return
	// }

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println("Migration Up Error:", err)
		return
	}

	r := &api.Router{}
	r.AddRoute("GET", "/post/", http.HandlerFunc(handlers.HandlePosts))
	r.AddRoute("POST", "/post/", http.HandlerFunc(handlers.CreateComment))
	r.AddRoute("POST", "/login", http.HandlerFunc(handlers.LoginHandler))
	r.AddRoute("POST", "/logout", http.HandlerFunc(handlers.LogoutHandler))
	r.AddRoute("GET", "/posts", http.HandlerFunc(handlers.HandlePosts))
	r.AddRoute("GET", "/", http.HandlerFunc(handlers.HomeHandler))
	r.AddRoute("GET", "/profile/", http.HandlerFunc(handlers.ProfileHandler))
	r.AddRoute("POST", "/", http.HandlerFunc(handlers.CreatePostHandler))
	r.AddRoute("POST", "/register", http.HandlerFunc(handlers.RegisterUser))
	r.AddRoute("GET", "/image/", http.HandlerFunc(handlers.GetImageHandler))
	r.AddRoute("GET", "/avatar/", http.HandlerFunc(handlers.GetAvaterFromUserId))
	r.AddRoute("GET", "/credential", http.HandlerFunc(handlers.GetCredential))
	r.AddRoute("POST", "/privacy", http.HandlerFunc(handlers.PrivacyHandler))
	r.AddRoute("GET", "/user/", http.HandlerFunc(handlers.GetNicknameFromId))
	r.AddRoute("GET", "/notifications", http.HandlerFunc(handlers.Notifications))
	r.AddRoute("GET", "/following", http.HandlerFunc(handlers.GetFollowing))
	r.AddRoute("GET", "/ws", http.HandlerFunc(handlers.WsHandler))

	// group-related handlers
	r.AddRoute("GET", "/fetch_memberships/", http.HandlerFunc(handlers.MembershipsHandler))
	r.AddRoute("GET", "/group/", http.HandlerFunc(handlers.GroupHandler))
	r.AddRoute("GET", "/groups", http.HandlerFunc(handlers.GroupsHandler))
	r.AddRoute("POST", "/create_group", http.HandlerFunc(handlers.CreateGroupHandler))
	r.AddRoute("POST", "/invite_member", http.HandlerFunc(handlers.InviteMemberHandler))
	r.AddRoute("POST", "/update_membership", http.HandlerFunc(handlers.UpdateMemberStatusHandler))
	r.AddRoute("POST", "/create_group_post", http.HandlerFunc(handlers.CreateGroupPostHandler))
	r.AddRoute("POST", "/fetch_your_requests", http.HandlerFunc(handlers.FetchYourRequestsHandler))
	r.AddRoute("GET", "/fetch_group_posts/", http.HandlerFunc(handlers.GroupPostsHandler))
	r.AddRoute("GET", "/fetch_group_post/", http.HandlerFunc(handlers.GroupPostHandler))
	r.AddRoute("GET", "/fetch_group_post_comment/", http.HandlerFunc(handlers.GroupPostCommentsHandler))
	r.AddRoute("POST", "/create_group_post_comment", http.HandlerFunc(handlers.CreateGroupPostCommentHandler))
	r.AddRoute("GET", "/fetch_group_events/", http.HandlerFunc(handlers.FetchGroupEventsHandler))
	r.AddRoute("POST", "/create_group_event", http.HandlerFunc(handlers.CreateGroupEventHandler))
	r.AddRoute("GET", "/fetch_user_event_status/", http.HandlerFunc(handlers.FetchUserEventStatusHandler))
	r.AddRoute("POST", "/update_event_status", http.HandlerFunc(handlers.UpdateEventStatusHandler))

	fmt.Println("Starting Go server")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Server Error:", err)
	}
}
