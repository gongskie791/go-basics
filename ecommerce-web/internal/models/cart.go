package models

type CartItem struct {
	ProductID int
	Quantity  int
}

type Cart struct {
	ID    int
	Items []CartItem
	Total float64
}
