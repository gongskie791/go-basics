package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"slices"
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
	mux.HandleFunc("/add", addHandler)
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Account []Account
		Count   int
	}{
		Account: account,
		Count:   len(account),
	}

	templates.ExecuteTemplate(w, "home.html", data)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	exists := slices.ContainsFunc(account, func(a Account) bool {
		return a.Username == username
	})

	if exists {
		fmt.Printf("Username %s used", username)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if username == "" {
		fmt.Print("Username is empty")
	}
	if password == "" {
		fmt.Print("Password is empty")
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
