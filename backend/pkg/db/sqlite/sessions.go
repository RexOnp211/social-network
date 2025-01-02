package db

import (
	"fmt"
	"log"
	"social-network/pkg/helpers"
	"time"
)

// SaveSession saves a new session in the database.
func SaveSession(token, nickname string, userID int, expiration time.Time) error {
	if DB == nil {
		return fmt.Errorf("db connection failed")
	}
	stmt, err := DB.Prepare(`INSERT INTO sessions (token, user_id, nickname, expires_at) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(token, userID, nickname, expiration)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSession deletes a session from the database.
func DeleteSession(token string) error {
	if DB == nil {
		return fmt.Errorf("db connection failed")
	}
	stmt, err := DB.Prepare(`DELETE FROM sessions WHERE token = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(token)
	if err != nil {
		return err
	}
	return nil
}

// GetSessionByUserID retrieves a session by user ID.
func GetSessionByUserID(userID int) (*helpers.Session, error) {
	var session helpers.Session
	err := DB.QueryRow(`SELECT token, user_id, expires_at FROM sessions WHERE nickname = ?`, userID).Scan(
		&session.SessionToken, &session.UserID, &session.ExpireTime)
	if err != nil {
		log.Printf("Error finding session for user ID %d: %v", userID, err)
		return nil, err
	}
	log.Printf("Session token found for user ID: %d", session.UserID)
	return &session, nil
}

// GetUserIDFromSession retrieves the user ID associated with a session token.
func GetUserIDFromSession(token string) (int, error) {
	if DB == nil {
		return 0, fmt.Errorf("db connection failed")
	}
	var userID int
	err := DB.QueryRow(`SELECT user_id FROM sessions WHERE token = ?`, token).Scan(&userID)
	if err != nil {
		log.Printf("Error finding session token: %v", err)
		return 0, err
	}
	log.Printf("Session token found for user ID: %d", userID)
	return userID, nil
}

func GetNicknameFromSession(token string) (string, error) {
	if DB == nil {
		return "", fmt.Errorf("db connection failed")
	}
	var Nickname string
	err := DB.QueryRow(`SELECT nickname FROM sessions WHERE token = ?`, token).Scan(&Nickname)
	if err != nil {
		log.Printf("Error finding session token: %v", err)
		return "", err
	}
	log.Printf("Session token found for user ID: %s", Nickname)
	return Nickname, nil
}

// ClearSessions deletes all sessions from the database.
func ClearSessions() error {
	if DB == nil {
		return fmt.Errorf("db connection failed")
	}
	_, err := DB.Exec("DELETE FROM sessions")
	if err != nil {
		return fmt.Errorf("failed to clear sessions: %v", err)
	}
	log.Println("All sessions have been cleared from the database")
	return nil
}

func ClearUserStatus() error {
	if DB == nil {
		return fmt.Errorf("db connection failed")
	}
	_, err := DB.Exec("DELETE FROM user_status")
	if err != nil {
		return fmt.Errorf("failed to clear user_status: %v", err)
	}
	log.Println("All user_status have been cleared from the database")
	return nil
}
