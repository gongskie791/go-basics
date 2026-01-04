package main

import (
	"log"
	"net/http"
	"session_with_database_authentication/internal/handlers"
	"session_with_database_authentication/internal/middlewares"
)

func setupRoutes() http.Handler {
	mux := http.NewServeMux()
	err := handlers.InitTemplate()
	if err != nil {
		log.Fatal("Failed to load templates:", err)
	}

	// Use mux.Handle and mux.HandleFunc instead!
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/dashboard", handlers.DashboardHandler)
	mux.HandleFunc("/login", handlers.Login)

	return middlewares.LoggingMiddleware(mux)
}
