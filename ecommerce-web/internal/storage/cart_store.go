package storage

import (
	"ecommerce-web/internal/models"
	"sync"
)

type CartStore struct {
	carts  map[int]models.Cart
	mu     sync.RWMutex
	nextID int
}

func NewCartStore() *CartStore {
	return &CartStore{
		carts:  make(map[int]models.Cart),
		nextID: 1,
	}
}

func (s *CartStore) CreateCart() models.Cart {
	s.mu.Lock()
	defer s.mu.Unlock()

	cart := models.Cart{
		ID:    s.nextID,
		Items: make([]models.CartItem, 0),
		Total: 0,
	}

	s.nextID++
	s.carts[cart.ID] = cart
	return cart
}

func (s *CartStore) GetCart(id int) (models.Cart, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock() // Fixed: was RLocker()

	cart, exists := s.carts[id]
	return cart, exists
}

func (s *CartStore) AddToStore(cartID int, productID int, quantity int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	cart, exist := s.carts[cartID]
	if !exist {
		return false
	}

	for i, item := range cart.Items {
		if item.ProductID == productID {
			cart.Items[i].Quantity += quantity
			s.carts[cartID] = cart
			return true
		}
	}

	cart.Items = append(cart.Items, models.CartItem{
		ProductID: productID,
		Quantity:  quantity,
	})
	s.carts[cartID] = cart
	return true
}

func (s *CartStore) ClearCart(cartID int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if cart, exists := s.carts[cartID]; exists {
		cart.Items = make([]models.CartItem, 0)
		cart.Total = 0
		s.carts[cartID] = cart
	}
}
