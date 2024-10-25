package handlers

import (
	"log"
	"net/http"
	"social-network/pkg/helpers"
	"social-network/pkg/db/sqlite"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
)

var sessions = map[string]helpers.Session{}

// NewSession creates a new session for the user
func NewSession(w http.ResponseWriter, username string, userID int) (string, error) {
	existingSession, err := db.GetSessionByUserID(userID)
	if err == nil && existingSession.SessionToken != "" {
		log.Println("Existing session found for user:", username)
		return existingSession.SessionToken, nil
	}

	token, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("Failed to generate UUID: %v", err)
		return "", err
	}

	session := helpers.Session{
		Username:     username,
		SessionToken: token.String(),
		ExpireTime:   time.Now().Add(100 * time.Minute),
		UserID:       userID,
	}
	log.Printf("Creating session for user %s with token %s", username, token.String())
	sessions[token.String()] = session

	expiration := time.Now().Add(4 * time.Hour)
	cookie := http.Cookie{
		Name:     "session_token",
		Value:    session.SessionToken,
		Expires:  expiration,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	err = db.SaveSession(token.String(), userID, expiration)
	if err != nil {
		log.Fatalf("Failed to save session to database: %v", err)
		return "", err
	}

	log.Printf("New session created for user %s with token %s", username, token.String())
	return token.String(), nil
}

func ValidateSession(w http.ResponseWriter, r *http.Request) string {
	token, err := getSessionToken(r)
	if err != nil {
		log.Println("Session token not found:", err)
		return ""
	}
	log.Printf("Session token found: %s", token)
	key, ok := sessions[token]
	if ok {
		log.Printf("Session valid for user: %s", key.Username)
		return key.Username
	} else {
		log.Printf("Invalid session token: %s", token)
		userID, err := db.GetUserIDFromSession(token)
		if err == nil {
			log.Printf("Valid session token found in database for user ID: %d", userID)
			return strconv.Itoa(userID)
		} else {
			log.Printf("Session token not found in database: %v", err)
		}
		return ""
	}
}

// SessionExpired checks if the session has expired
func SessionExpired(w http.ResponseWriter, r *http.Request) bool {
	token, err := getSessionToken(r)
	if err != nil {
		log.Println("Session token not found:", err)
		return true
	}
	log.Printf("Session token for expiration check: %s", token)
	key, ok := sessions[token]
	if ok {
		isExpired := key.ExpireTime.Before(time.Now())
		log.Printf("Session for user %s expired: %v", key.Username, isExpired)
		return isExpired
	}
	return true
}

// CloseSession closes the session and deletes the cookie
func CloseSession(w http.ResponseWriter, r *http.Request) {
	token, err := getSessionToken(r)
	if err != nil {
		http.Error(w, "No session token found", http.StatusBadRequest)
		return
	}
	log.Printf("Attempting to close session: %s", token)
	_, ok := sessions[token]
	if ok {
		delete(sessions, token)
		cookie := http.Cookie{
			Name:   "session_token",
			Value:  "",
			MaxAge: -1,
			Path:   "/",
		}
		http.SetCookie(w, &cookie)
		err = db.DeleteSession(token)
		if err != nil {
			log.Printf("Failed to delete session from database: %v", err)
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logout successful"))
}

// RequireLogin is a middleware that checks for a valid session
func RequireLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if SessionExpired(w, r) {
			http.Error(w, "Session expired", http.StatusUnauthorized)
			return
		}
		username := ValidateSession(w, r)
		if username == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// func getUserIDFromSession(r *http.Request) (int, error) {
// 	cookie, err := r.Cookie("session_token")
// 	if err != nil {
// 		log.Println("No session cookie found:", err)
// 		return 0, err
// 	}
// 	sessionToken := cookie.Value
// 	log.Println("Session token received:", sessionToken)
// 	return db.GetUserIDFromSession(sessionToken)
// }

// getSessionToken extracts the session token from the request
func getSessionToken(r *http.Request) (string, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}
