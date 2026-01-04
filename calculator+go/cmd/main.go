package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"slices"
	"strings"
)

type Account struct {
	ID       int64
	Username string
	Password string
}

var account []Account
var nextID int64 = 1

var templates = template.Must(template.ParseFiles(
	"templates/layout.html",
	"templates/home.html",
	// "templates/task.html",
))

func main() {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	account = append(account, Account{
		ID:       nextID,
		Username: "mark",
		Password: "gongkskie",
	})
	nextID++

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/account/new", addHandler)
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Account    []Account
		Count      int
		ActivePage string // Add this
	}{
		Account:    account,
		Count:      len(account),
		ActivePage: "home", // Set active page
	}

	templates.ExecuteTemplate(w, "home.html", data)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	username := template.HTMLEscapeString(r.FormValue("username"))
	password := template.HTMLEscapeString(r.FormValue("password"))

	exists := slices.ContainsFunc(account, func(a Account) bool {
		return a.Username == username
	})

	if exists {
		fmt.Printf("Username %s used", username)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if strings.TrimSpace(username) == "" || strings.TrimSpace(password) == "" {
		fmt.Print("username and password are required")
	}

	newAccount := Account{
		ID:       nextID,
		Username: username,
		Password: password,
	}
	account = append(account, newAccount)
	nextID++

	fmt.Print("\n Accounts: ", account)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
