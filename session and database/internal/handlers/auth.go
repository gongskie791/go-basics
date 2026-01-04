package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"sync"
	"time"
)

var (
	sessions = make(map[string]Session)
	mu       sync.RWMutex
)

type Session struct {
	UserID    int64
	Username  string
	ExpiresAt time.Time
}

var nextId int64 = 1

func GenerateSessionID() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// if user is already logged in, redirect to dashboard
	_, ok := GetSession(r)

	if ok {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	// if not logged in
	if r.Method != http.MethodPost {
		http.ServeFile(w, r, "web/templates/pages/login.html")
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username != "admin" || password != "secret" {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	SessionID := GenerateSessionID()
	mu.Lock()
	sessions[SessionID] = Session{
		UserID:    nextId,
		Username:  username,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	nextId++
	mu.Unlock()

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    SessionID,
		Path:     "/",
		MaxAge:   86400,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func GetSession(r *http.Request) (Session, bool) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return Session{}, false
	}

	mu.RLock()
	session, exists := sessions[cookie.Value]
	mu.RUnlock()

	if !exists || time.Now().After(session.ExpiresAt) {
		return Session{}, false
	}

	return session, true
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err == nil {
		mu.Lock()
		delete(sessions, cookie.Value)
		mu.Unlock()
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
