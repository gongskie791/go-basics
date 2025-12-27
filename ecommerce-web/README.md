# E-Commerce Web Application

A Go web application built with the standard library. This is a web version of a console-based e-commerce system.

## Project Structure

```
ecommerce-web/
├── cmd/
│   └── web/
│       └── main.go           # Entry point - starts the web server
├── internal/
│   ├── models/               # Data structures (Product, Cart, CartItem)
│   │   ├── product.go
│   │   └── cart.go
│   ├── services/             # Business logic (same as console app)
│   │   ├── product_service.go
│   │   └── cart_service.go
│   ├── storage/              # Data storage (in-memory)
│   │   ├── product_store.go
│   │   └── cart_store.go
│   └── handlers/             # HTTP handlers (NEW - replaces console menu)
│       └── handlers.go
├── templates/                # HTML templates
│   ├── layout.html           # Base template (nav, footer)
│   ├── home.html             # Product listing
│   ├── cart.html             # Shopping cart
│   └── new_product.html      # Add product form
├── static/
│   └── css/
│       └── style.css         # Styling
├── go.mod
└── README.md
```

## How It Maps to Your Console App

| Console Menu Option | Web Route | Handler Function |
|---------------------|-----------|------------------|
| 1. Display Products | GET /     | Home()           |
| 2. Add to Cart      | POST /cart/add | AddToCart() |
| 3. View Cart        | GET /cart | ViewCart()       |
| 4. Checkout         | POST /cart/checkout | Checkout() |
| 5. Add New Product  | GET /products/new + POST /products | NewProduct() + CreateProduct() |

## Running the Application

1. Make sure you have Go installed (1.21 or later)

2. Navigate to the project directory:
   ```bash
   cd ecommerce-web
   ```

3. Run the application:
   ```bash
   go run ./cmd/web/
   ```

4. Open your browser and go to: http://localhost:8080

## Key Concepts

### 1. HTTP Handler
A handler is a function that receives HTTP requests and sends responses. It replaces your console menu's switch cases:

```go
// Console version
case "1":
    displayProducts(productService)

// Web version
func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
    products := h.productService.ListProducts()
    h.templates.ExecuteTemplate(w, "home.html", products)
}
```

### 2. Routes
Routes map URLs to handler functions:

```go
mux.HandleFunc("GET /", h.Home)           // When user visits /
mux.HandleFunc("POST /cart/add", h.AddToCart)  // When form is submitted
```

### 3. Templates
HTML templates with Go syntax for dynamic content:

```html
{{range .Products}}
    <h3>{{.Name}}</h3>
    <p>${{.Price}}</p>
{{end}}
```

### 4. Form Handling
Forms submit data via POST requests:

```html
<form action="/cart/add" method="POST">
    <input type="hidden" name="product_id" value="{{.ID}}">
    <button type="submit">Add to Cart</button>
</form>
```

In the handler:
```go
r.ParseForm()
productID := r.FormValue("product_id")
```

## Next Steps

1. **Add HTMX** - Make the page interactive without full reloads
2. **Add a database** - Replace in-memory storage with SQLite or PostgreSQL
3. **Add user authentication** - Login/signup functionality
4. **Add sessions** - Each user gets their own cart

## Learning Resources

- [Go net/http package](https://pkg.go.dev/net/http)
- [Go html/template package](https://pkg.go.dev/html/template)
- [HTMX](https://htmx.org/) - For adding interactivity later
