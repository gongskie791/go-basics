package storage

import (
	"ecommerce-web/internal/models"
	"sync"
)

type ProductStore struct {
	products []models.Product
	mu       sync.RWMutex
	nextID   int
}

func NewProductStore() *ProductStore {
	return &ProductStore{
		products: make([]models.Product, 0),
		nextID:   1,
	}
}

func (s *ProductStore) Add(product models.Product) models.Product {
	s.mu.Lock()
	defer s.mu.Unlock()

	product.ID = s.nextID
	s.nextID++
	s.products = append(s.products, product)
	return product
}

func (s *ProductStore) GetAll() []models.Product {
	s.mu.RLock()
	defer s.mu.RUnlock()

	results := make([]models.Product, len(s.products))
	copy(results, s.products)
	return results
}

func (s *ProductStore) GetByID(id int) (models.Product, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, p := range s.products {
		if p.ID == id {
			return p, true
		}
	}
	return models.Product{}, false
}

func (s *ProductStore) UpdateStock(id int, quantity int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range s.products {
		if s.products[i].ID == id {
			if s.products[i].Stock >= quantity {
				s.products[i].Stock -= quantity
				return true
			}
			return false
		}
	}
	return false
}
