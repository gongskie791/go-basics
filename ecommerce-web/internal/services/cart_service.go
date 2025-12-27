package services

import (
	"ecommerce-web/internal/models"
	"ecommerce-web/internal/storage"
)

type CartService struct {
	cartStore     *storage.CartStore
	productStore  *storage.ProductStore
	currentCartID int
}

func NewCartService(cartStore *storage.CartStore, productStore *storage.ProductStore) *CartService {
	cs := &CartService{
		cartStore:    cartStore,
		productStore: productStore,
	}

	cart := cartStore.CreateCart()
	cs.currentCartID = cart.ID
	return cs
}

func (s *CartService) AddItem(productID int, quantity int) (bool, string) {
	product, exists := s.productStore.GetByID(productID)
	if !exists {
		return false, "Product not found"
	}

	if product.Stock < quantity {
		return false, "Insufficient stock"
	}

	success := s.cartStore.AddToStore(s.currentCartID, productID, quantity)
	if success {
		return true, "Added to cart successfully" // Fixed: was returning false
	}

	return false, "Failed to add to cart"
}

func (s *CartService) GetCart() (models.Cart, float64, []models.Product) {
	cart, _ := s.cartStore.GetCart(s.currentCartID)

	var total float64
	var items []models.Product

	for _, item := range cart.Items {
		product, exists := s.productStore.GetByID(item.ProductID)
		if exists {
			items = append(items, product)
			total += product.Price * float64(item.Quantity)
		}
	}

	return cart, total, items
}

func (s *CartService) GetCartItemCount() int {
	cart, _, _ := s.GetCart()
	count := 0
	for _, item := range cart.Items {
		count += item.Quantity
	}
	return count
}

func (s *CartService) Checkout() (bool, string, float64) {
	cart, total, items := s.GetCart()

	if len(items) == 0 {
		return false, "Cart is empty", 0
	}

	// Update stock
	for _, item := range cart.Items {
		s.productStore.UpdateStock(item.ProductID, item.Quantity)
	}

	// Clear cart
	s.cartStore.ClearCart(s.currentCartID)

	// Create new cart for next session
	newCart := s.cartStore.CreateCart()
	s.currentCartID = newCart.ID

	return true, "Checkout successful!", total
}
