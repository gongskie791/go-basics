package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"ecommerce-web/internal/models"
	"ecommerce-web/internal/services"
)

// Handler holds our services (like dependency injection)
type Handler struct {
	productService *services.ProductService
	cartService    *services.CartService
	templates      *template.Template
}

// NewHandler creates a handler with all templates loaded
func NewHandler(ps *services.ProductService, cs *services.CartService) *Handler {
	// Define custom template functions
	funcMap := template.FuncMap{
		"mul": func(a float64, b float64) float64 {
			return a * b
		},
		"toFloat": func(i int) float64 {
			return float64(i)
		},
	}

	// Load all templates at startup with custom functions
	tmpl := template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/*.html"))

	return &Handler{
		productService: ps,
		cartService:    cs,
		templates:      tmpl,
	}
}

// PageData is passed to templates (like what you printed in console)
type PageData struct {
	Title     string
	Products  []models.Product
	Cart      models.Cart
	CartItems []models.Product
	CartTotal float64
	CartCount int
	Message   string
	Error     string
}

// Home displays all products (replaces case "1": displayProducts)
func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	products := h.productService.ListProducts()

	data := PageData{
		Title:     "Products",
		Products:  products,
		CartCount: h.cartService.GetCartItemCount(),
		Message:   r.URL.Query().Get("message"),
		Error:     r.URL.Query().Get("error"),
	}

	h.templates.ExecuteTemplate(w, "home.html", data)
}

// ViewCart shows the cart (replaces case "3": viewCart)
func (h *Handler) ViewCart(w http.ResponseWriter, r *http.Request) {
	cart, total, items := h.cartService.GetCart()

	data := PageData{
		Title:     "Your Cart",
		Cart:      cart,
		CartItems: items,
		CartTotal: total,
		CartCount: h.cartService.GetCartItemCount(),
		Message:   r.URL.Query().Get("message"),
		Error:     r.URL.Query().Get("error"),
	}

	h.templates.ExecuteTemplate(w, "cart.html", data)
}

// AddToCart handles form submission (replaces case "2": addToCart)
func (h *Handler) AddToCart(w http.ResponseWriter, r *http.Request) {
	// Parse form data (like scanner.Scan() in console)
	r.ParseForm()

	productID, err := strconv.Atoi(r.FormValue("product_id"))
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil || quantity < 1 {
		quantity = 1
	}

	success, message := h.cartService.AddItem(productID, quantity)

	// Redirect back to home with a message
	if success {
		http.Redirect(w, r, "/?message="+message, http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/?error="+message, http.StatusSeeOther)
	}
}

// Checkout processes the order (replaces case "4": checkout)
func (h *Handler) Checkout(w http.ResponseWriter, r *http.Request) {
	success, message, total := h.cartService.Checkout()

	if success {
		// Show success with total
		http.Redirect(w, r, "/cart?message="+message+" Total: $"+strconv.FormatFloat(total, 'f', 2, 64), http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/cart?error="+message, http.StatusSeeOther)
	}
}

// NewProduct shows the add product form (replaces case "5" prompt)
func (h *Handler) NewProduct(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:     "Add New Product",
		CartCount: h.cartService.GetCartItemCount(),
	}

	h.templates.ExecuteTemplate(w, "new_product.html", data)
}

// CreateProduct handles the form submission (replaces case "5": addNewProduct)
func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := r.FormValue("name")
	description := r.FormValue("description")

	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		http.Redirect(w, r, "/products/new?error=Invalid price", http.StatusSeeOther)
		return
	}

	stock, err := strconv.Atoi(r.FormValue("stock"))
	if err != nil {
		http.Redirect(w, r, "/products/new?error=Invalid stock", http.StatusSeeOther)
		return
	}

	h.productService.AddProduct(name, description, price, stock)

	http.Redirect(w, r, "/?message=Product added successfully", http.StatusSeeOther)
}
