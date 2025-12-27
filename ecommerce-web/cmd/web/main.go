package main

import (
	"log"
	"net/http"

	"ecommerce-web/internal/handlers"
	"ecommerce-web/internal/services"
	"ecommerce-web/internal/storage"
)

func main() {
	// Initialize storage (same as your console app)
	productStore := storage.NewProductStore()
	cartStore := storage.NewCartStore()

	// Initialize services (same as your console app)
	productService := services.NewProductService(productStore)
	cartService := services.NewCartService(cartStore, productStore)

	// Add some sample products to start with
	productService.AddProduct("Laptop", "High-performance laptop", 999.99, 10)
	productService.AddProduct("Mouse", "Wireless mouse", 29.99, 50)
	productService.AddProduct("Keyboard", "Mechanical keyboard", 79.99, 30)
	productService.AddProduct("Monitor", "27-inch 4K display", 399.99, 15)

	// Initialize handlers (this is new - replaces your console menu)
	h := handlers.NewHandler(productService, cartService)

	// Set up routes (maps URLs to functions, like your switch statement)
	mux := http.NewServeMux()
	// Test route
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, it works!"))
	})

	// Serve static files (CSS, JS, images)
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Page routes
	mux.HandleFunc("GET /{$}", h.Home)                // Display products (case "1")
	mux.HandleFunc("GET /cart", h.ViewCart)           // View cart (case "3")
	mux.HandleFunc("GET /products/new", h.NewProduct) // Add product form (case "5")

	// Action routes (form submissions)
	mux.HandleFunc("POST /cart/add", h.AddToCart)     // Add to cart (case "2")
	mux.HandleFunc("POST /cart/checkout", h.Checkout) // Checkout (case "4")
	mux.HandleFunc("POST /products", h.CreateProduct) // Create product (case "5")

	// Start server
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
