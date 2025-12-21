package services

import (
	"ecommerce/models"
	"ecommerce/storage"
)

type ProductService struct {
	store *storage.ProductStore
}

func NewProductService(store *storage.ProductStore) *ProductService {
	return &ProductService{store: store}
}

func (s *ProductService) AddProduct(name, description string, price float64, stock int) models.Product {
	product := models.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
	}

	return s.store.Add(product)
}

func (s *ProductService) ListProducts() []models.Product {
	return s.store.GetAll()
}

func (s *ProductService) GetProduct(id int) (models.Product, bool) {
	return s.store.GetByID(id)
}

func (s *ProductService) UpdateStock(id int, quantity int) bool {
	return s.store.UpdateStock(id, quantity)
}
