// go-na-pratica/Scripts/products/product.go
package products // Este arquivo pertence ao pacote 'products'

// Product representa um produto com ID, Nome e Preço.
// Começa com letra maiúscula para ser exportável e acessível por outros pacotes.
type Product struct {
	ID    string
	Name  string
	Price float64
}
