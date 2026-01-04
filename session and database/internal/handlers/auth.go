package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"html/template"
	"net/http"
	"sync"
	"time"
)

var (
	sessions     = make(map[string]Session)
	mu           sync.RWMutex
	loginTmpl    = template.Must(template.ParseFiles("web/templates/pages/login.html"))    // ← Add this
	registerTmpl = template.Must(template.ParseFiles("web/templates/pages/register.html")) // ← Add this
	users        = make(map[string]string)
	usersMu      sync.RWMutex
)

type LoginPageData struct {
	Error string
}

type RegisterPageData struct {
	Error string
}

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

	if r.Method != http.MethodPost {
		loginTmpl.Execute(w, LoginPageData{})
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	usersMu.RLock()
	storedPassword, exists := users[username]
	usersMu.RUnlock()

	if !exists || storedPassword != password {
		w.WriteHeader(http.StatusUnauthorized)
		loginTmpl.Execute(w, LoginPageData{Error: "Invalid username or password"})
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

func Register(w http.ResponseWriter, r *http.Request) {
	_, ok := GetSession(r)

	if ok {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	if r.Method != http.MethodPost {
		registerTmpl.Execute(w, RegisterPageData{})
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")

	// validation
	if username == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		registerTmpl.Execute(w, RegisterPageData{Error: "Username and password are required"})
		return
	}

	if password != confirmPassword {
		w.WriteHeader(http.StatusBadRequest)
		registerTmpl.Execute(w, RegisterPageData{Error: "Password do not match"})
		return
	}

	usersMu.RLock()
	_, exist := users[username]
	usersMu.RUnlock()

	if exist {
		w.WriteHeader(http.StatusConflict)
		registerTmpl.Execute(w, RegisterPageData{Error: "Username already taken"})
		return
	}

	// Create user
	usersMu.Lock()
	users[username] = password // In real app: hash the password!
	usersMu.Unlock()

	// Redirect to login
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func GetSession(r *http.Request) (Session, bool) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return Session{}, false
	}

	mu.RLock()
	session, exists := sessions[cookie.Value]
	mu.RUnlock()

	fmt.Print(session)
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
