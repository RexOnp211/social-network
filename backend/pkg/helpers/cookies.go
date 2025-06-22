package helpers

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func SetCookie(w http.ResponseWriter, value string) {
	cookie := http.Cookie{
		Name:     "session_id",
		Value:    value,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}
	feDomain := os.Getenv("FE_DOMAIN")
	if feDomain != "" {
		cookie.Domain = feDomain // ex: "social-network-frontend-4ub2.onrender.com"
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
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}
	feDomain := os.Getenv("FE_DOMAIN")
	if feDomain != "" {
		cookie.Domain = feDomain
	}
	http.SetCookie(w, cookie)
}
