package models

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Stock       int
}

type CartItem struct {
	ProductID int
	Quantity  int
}

type Cart struct {
	ID    int
	Items []CartItem
	Total float64
}
