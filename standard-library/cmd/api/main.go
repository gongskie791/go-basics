package main

import (
	"first/project/pkg/recipes"
	"net/http"
)

func main() {
	store := recipes.NewMemStore()
	recipesHandler := NewRecipesHandler(store)

	mux := http.NewServeMux()
	mux.Handle("/recipes", recipesHandler)
	mux.Handle("/recipes/", recipesHandler)

	http.ListenAndServe(":8080", mux)

}
