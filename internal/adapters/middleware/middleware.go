package middleware

import (
	"leetboard/internal/core/util"
	"log/slog"
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
		slog.Info("Cookie not found")
		CreateSessionCookie(w)
		return
	}
}

func CreateSessionCookie(w http.ResponseWriter) {
	userSessionID := util.UserSessionGenerator()
	slog.Debug("Creating session cookie...", "user session ID", userSessionID)
	sessionCookie := http.Cookie{Name: "session_id", Value: userSessionID, Expires: time.Now().Add(100 * time.Second)}
	http.SetCookie(w, &sessionCookie)
}
