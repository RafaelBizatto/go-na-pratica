package main

import (
	"fmt"
	"go-na-pratica/Scripts/products" // Importa o pacote 'products' do seu módulo
)

func main() {
	fmt.Println("Iniciando o Gerenciamento de Produtos Go (Projeto Completo)!")

	// Adicionando alguns produtos
	products.AddProduct("P001", "Smartphone X", 1200.00)
	products.AddProduct("P002", "Smart TV 50", 2500.00)
	products.AddProduct("P003", "Fone de Ouvido Bluetooth", 150.00)

	products.ListProducts()

	// --- Testando a busca por ID ---
	fmt.Println("\n--- Testando Busca de Produto ---")
	productIDToFind := "P002"
	foundProduct, err := products.FindProductByID(productIDToFind)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Produto encontrado: ID: %s, Nome: %s, Preço: %.2f\n", foundProduct.ID, foundProduct.Name, foundProduct.Price)
	}

	productIDToFind = "P999" // Testando um ID que não existe
	foundProduct, err = products.FindProductByID(productIDToFind)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Produto encontrado: ID: %s, Nome: %s, Preço: %.2f\n", foundProduct.ID, foundProduct.Name, foundProduct.Price)
	}

	// --- Testando a atualização de preço ---
	fmt.Println("\n--- Testando Atualização de Preço ---")
	err = products.UpdateProductPrice("P001", 1150.00) // Atualiza o preço do Smartphone X
	if err != nil {
		fmt.Println(err)
	}
	products.ListProducts() // Lista para ver a mudança

	err = products.UpdateProductPrice("P003", -10.00) // Tenta atualizar com preço negativo
	if err != nil {
		fmt.Println(err)
	}

	// --- Testando a remoção por ID ---
	fmt.Println("\n--- Testando Remoção de Produto ---")
	err = products.RemoveProductByID("P002") // Remove a Smart TV
	if err != nil {
		fmt.Println(err)
	}
	products.ListProducts() // Lista para ver a remoção

	err = products.RemoveProductByID("P999") // Tenta remover um ID que não existe
	if err != nil {
		fmt.Println(err)
	}
	products.ListProducts() // Lista novamente

	fmt.Println("\nPrograma Go finalizado.")
}
