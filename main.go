package main

import "fmt"

// Product representa um produto com ID, Nome e Preço.
type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// products é um slice global para armazenar todos os produtos.
var products []Product

// AddProduct adiciona um novo produto à lista.
func AddProduct(id, name string, price float64) {
	newProduct := Product{
		ID:    id,
		Name:  name,
		Price: price,
	}
	products = append(products, newProduct)
	fmt.Printf("Produto \"%s\" (ID: %s) adicionado com sucesso!\n", name, id)
}

// ListProducts exibe todos os produtos cadastrados.
func ListProducts() {
	fmt.Println("\n--- Lista de Produtos ---")
	if len(products) == 0 {
		fmt.Println("Nenhum produto cadastrado ainda.")
		return
	}
	for i, p := range products {
		fmt.Printf("%d. ID: %s, Nome: %s, Preço: %.2f\n", i+1, p.ID, p.Name, p.Price)
	}
	fmt.Println("-------------------------")
}

func main() {
	fmt.Println("Iniciando o Gerenciamento de Produtos Go!")

	// Adicionando alguns produtos
	AddProduct("P001", "Smartphone X", 1200.00)
	AddProduct("P002", "Smart TV 50", 2500.00)
	AddProduct("P003", "Fone de Ouvido Bluetooth", 150.00)

	// Listando os produtos
	ListProducts()

	// Tentando adicionar mais um produto
	AddProduct("P004", "Teclado Mecânico", 300.00)

	// Listando novamente para ver o novo produto
	ListProducts()
}
