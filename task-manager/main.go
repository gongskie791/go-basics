package main

import (
	"html/template"
	"net/http"
	"strconv"
)

// Task represents a single to-do item
type Task struct {
	ID        int
	Title     string
	Completed bool
}

// In-memory storage (resets when server restarts)
var tasks []Task
var nextID = 1

// Parse templates once at startup
var templates = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	// Define routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/toggle", toggleHandler)
	http.HandleFunc("/delete", deleteHandler)

	// Start server
	println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// homeHandler displays all tasks
func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Tasks []Task
		Count int
	}{
		Tasks: tasks,
		Count: len(tasks),
	}
	templates.Execute(w, data)
}

// addHandler creates a new task
func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	title := r.FormValue("title")
	if title != "" {
		tasks = append(tasks, Task{
			ID:        nextID,
			Title:     title,
			Completed: false,
		})
		nextID++
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// toggleHandler marks a task complete/incomplete
func toggleHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = !tasks[i].Completed
			break
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// deleteHandler removes a task
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
