package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"ecommerce-web/internal/models"
	"ecommerce-web/internal/services"
)

type Handler struct {
	productService *services.ProductService
	cartService    *services.CartService
}

func NewHandler(ps *services.ProductService, cs *services.CartService) *Handler {
	return &Handler{
		productService: ps,
		cartService:    cs,
	}
}

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

// Home - simple test
func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), 500)
		return
	}
	tmpl.Execute(w, nil)
}

// ViewCart - simple test
func (h *Handler) ViewCart(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/cart.html")
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), 500)
		return
	}
	tmpl.Execute(w, nil)
}

func (h *Handler) AddToCart(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	productID, _ := strconv.Atoi(r.FormValue("product_id"))
	quantity, _ := strconv.Atoi(r.FormValue("quantity"))
	if quantity < 1 {
		quantity = 1
	}
	h.cartService.AddItem(productID, quantity)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) Checkout(w http.ResponseWriter, r *http.Request) {
	h.cartService.Checkout()
	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

func (h *Handler) NewProduct(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/new_product.html")
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), 500)
		return
	}
	tmpl.Execute(w, nil)
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	description := r.FormValue("description")
	price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
	stock, _ := strconv.Atoi(r.FormValue("stock"))
	h.productService.AddProduct(name, description, price, stock)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
