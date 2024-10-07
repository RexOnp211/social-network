package pkg

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func SetCookie(w http.ResponseWriter, value string) {
	domain := "localhost:3000"
	if os.Getenv("FE_URL") != "" {
		domain = os.Getenv("FE_URL")
	}
	cookie := http.Cookie{
		Name:     "session_Id",
		Value:    value,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour),
		Domain:   domain,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}

func GetCookie(r *http.Request, userId string) string {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		fmt.Println("Cookie does not exist")
		return ""
	}
	return cookie.Value
}

func DeleteCookie(w http.ResponseWriter, r *http.Request) {
	domain := "localhost:3000"
	if os.Getenv("FE_URL") != "" {
		domain = os.Getenv("FE_URL")
	}
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   "",
		Path:    "/",
		Domain:  domain,
		Expires: time.Unix(0, 0),
	}
	http.SetCookie(w, cookie)
}
