package middleware

import (
	"fmt"
	"leetboard/internal/core/util"
	"net/http"
	"time"
)

type Session struct {
	ID string
}


// COOKIEs
func CheckCookie(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("session_id")
	if err != nil {
		fmt.Println("cookie not found")
		CreateSessionCookie(w)
		return
	}
}

func CreateSessionCookie(w http.ResponseWriter) {
	fmt.Println("creating session cookie...")
	sessionCookie := http.Cookie{Name: "session_id", Value: util.UserSessionGenerator(), Expires: time.Now().Add(100 * time.Second)}
	http.SetCookie(w, &sessionCookie)
}
