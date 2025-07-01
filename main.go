package main

import (
	"fmt"
)

func main() {
	fmt.Println("Iniciando o Gerenciamento de Produtos Go (Modularizado)!")

	// Adicionando alguns produtos
	AddProduct("P001", "Smartphone X", 1200.00)
	AddProduct("P002", "Smart TV 50", 2500.00)
	AddProduct("P003", "Fone de Ouvido Bluetooth", 150.00)

	ListProducts()

	// --- Testando a busca por ID ---
	fmt.Println("\n--- Testando Busca de Produto ---")
	productIDToFind := "P002"
	foundProduct, err := FindProductByID(productIDToFind)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Produto encontrado: ID: %s, Nome: %s, Preço: %.2f\n", foundProduct.ID, foundProduct.Name, foundProduct.Price)
	}

	productIDToFind = "P999"
	foundProduct, err = FindProductByID(productIDToFind)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Produto encontrado: ID: %s, Nome: %s, Preço: %.2f\n", foundProduct.ID, foundProduct.Name, foundProduct.Price)
	}

	// --- Testando a atualização de preço ---
	fmt.Println("\n--- Testando Atualização de Preço ---")
	err = UpdateProductPrice("P001", 1150.00)
	if err != nil {
		fmt.Println(err)
	}
	ListProducts()

	err = UpdateProductPrice("P003", -10.00)
	if err != nil {
		fmt.Println(err)
	}

	// --- Testando a remoção por ID ---
	fmt.Println("\n--- Testando Remoção de Produto ---")
	err = RemoveProductByID("P002")
	if err != nil {
		fmt.Println(err)
	}
	ListProducts()

	err = RemoveProductByID("P999")
	if err != nil {
		fmt.Println(err)
	}
	ListProducts()

	fmt.Println("\nPrograma finalizado.")
}
