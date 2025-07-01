package main

// Product representa um produto com ID, Nome e Preço.
type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
