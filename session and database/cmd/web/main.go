package main

import (
	"fmt"
	"log"
	"net/http"
	"session_with_database_authentication/internal/config"
)

func main() {
	fmt.Print("Lezgow")
	/* db := */ config.DB()

	handler := setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", handler)) // Use the handler!
}
