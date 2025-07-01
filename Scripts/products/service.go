package products

import "fmt" // Precisamos importar fmt aqui para usar fmt.Printf e fmt.Errorf

// products é um slice global dentro DO PACOTE products.
var products []Product

// AddProduct adiciona um novo produto à lista.
// Funções exportáveis começam com letra maiúscula.
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

// FindProductByID busca um produto pelo seu ID.
func FindProductByID(id string) (*Product, error) {
	for i, p := range products {
		if p.ID == id {
			return &products[i], nil
		}
	}
	return nil, fmt.Errorf("produto com ID %s não encontrado", id)
}

// RemoveProductByID remove um produto da lista pelo seu ID.
func RemoveProductByID(id string) error {
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			fmt.Printf("Produto com ID %s removido com sucesso!\n", id)
			return nil
		}
	}
	return fmt.Errorf("produto com ID %s não encontrado para remoção", id)
}

// UpdateProductPrice atualiza o preço de um produto existente.
func UpdateProductPrice(id string, newPrice float64) error {
	if newPrice < 0 {
		return fmt.Errorf("o novo preço não pode ser negativo")
	}

	for i := range products {
		if products[i].ID == id {
			products[i].Price = newPrice
			fmt.Printf("Preço do produto com ID %s atualizado para %.2f!\n", id, newPrice)
			return nil
		}
	}
	return fmt.Errorf("produto com ID %s não encontrado para atualização", id)
}
