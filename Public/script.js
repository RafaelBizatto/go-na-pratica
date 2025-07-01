// script.js

// Array para simular nossos produtos (substituindo o slice 'products' do Go)
let products = [];

// Elementos HTML que vamos manipular
const productsList = document.getElementById('productsList');
const noProductsMessage = document.getElementById('noProductsMessage');
const addProductForm = document.getElementById('addProductForm');

// Função para exibir os produtos na interface
function displayProducts() {
    productsList.innerHTML = ''; // Limpa a lista atual antes de renderizar novamente

    if (products.length === 0) {
        noProductsMessage.style.display = 'block'; // Mostra a mensagem de "nenhum produto"
    } else {
        noProductsMessage.style.display = 'none'; // Esconde a mensagem
        products.forEach(product => {
            const listItem = document.createElement('li'); // Cria um novo item de lista (<li>)
            listItem.innerHTML = `
                <span><strong>ID:</strong> ${product.id} | <strong>Nome:</strong> ${product.name} | <strong>Preço:</strong> R$ ${product.price.toFixed(2)}</span>
            `;
            productsList.appendChild(listItem); // Adiciona o item à lista UL
        });
    }
}

// Função para adicionar um novo produto (chamada pelo formulário)
function addProduct(event) {
    event.preventDefault(); // Impede o comportamento padrão de recarregar a página ao enviar o formulário

    const id = document.getElementById('productId').value;
    const name = document.getElementById('productName').value;
    const price = parseFloat(document.getElementById('productPrice').value); // Converte para número

    // Validação básica
    if (!id || !name || isNaN(price) || price < 0) {
        alert('Por favor, preencha todos os campos corretamente. O preço não pode ser negativo.');
        return;
    }

    // Verifica se o ID já existe
    const existingProduct = products.find(p => p.id === id);
    if (existingProduct) {
        alert(`Produto com ID "${id}" já existe. Por favor, use um ID único.`);
        return;
    }

    const newProduct = {
        id: id,
        name: name,
        price: price
    };

    products.push(newProduct); // Adiciona o novo produto ao array
    alert(`Produto "${name}" adicionado com sucesso!`);

    addProductForm.reset(); // Limpa os campos do formulário
    displayProducts(); // Atualiza a lista exibida
}

// Event Listeners:
// Quando a página é carregada, exibe os produtos iniciais (se houver)
document.addEventListener('DOMContentLoaded', displayProducts);

// Quando o formulário de adicionar produto é enviado, chama a função addProduct
addProductForm.addEventListener('submit', addProduct);


// --- Adicionando alguns produtos iniciais (como o nosso main.go) ---
// Você pode descomentar e adicionar produtos aqui para testar
// products.push({ id: "P001", name: "Smartphone X", price: 1200.00 });
// products.push({ id: "P002", name: "Smart TV 50", price: 2500.00 });
// products.push({ id: "P003", name: "Fone de Ouvido Bluetooth", price: 150.00 });
// products.push({ id: "P004", name: "Teclado Mecânico", price: 300.00 });
// Ao iniciar com produtos, a função displayProducts() no DOMContentLoaded já os mostrará.