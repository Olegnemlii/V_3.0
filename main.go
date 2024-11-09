package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var products = []Product{
	{ID: 1, Name: "Product 1"},
	{ID: 2, Name: "Product 2"},
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	// Возвращает все продукты
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	// Обновляет товар
	var updateProduct Product
	json.NewDecoder(r.Body).Decode(&updateProduct)
	for i, p := range products {
		if p.ID == updateProduct.ID {
			products[i] = updateProduct
			fmt.Fprint(w, "Product updated")
			return
		}
	}

	fmt.Fprint(w, "Product not found")
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	// Удатяет продукт
	var productID int
	json.NewDecoder(r.Body).Decode(&productID)

	for i, p := range products {
		if p.ID == productID {
			products = append(products[:i], products[i+1:]...)
			fmt.Fprint(w, "Product deleted")
			return
		}
	}

	fmt.Fprint(w, "Product not found")
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	// Добавляет новый продукт
	var newProduct Product
	json.NewDecoder(r.Body).Decode(&newProduct)

	products = append(products, newProduct)
	fmt.Fprint(w, "Product created")
}

func getProductByID(w http.ResponseWriter, r *http.Request) {
	// Возвращает конкретный продукт по id
	path := r.URL.Path
	parts := strings.Split(path, "/")
	productID := parts[len(parts)-1]
	fmt.Fprintf(w, "Product ID: %s", productID)
}

func main() {
	http.HandleFunc("/products", getAllProducts)
	http.HandleFunc("/products/update", updateProduct)
	http.HandleFunc("/products/delete", deleteProduct)
	http.HandleFunc("/products/create", createProduct)
	http.HandleFunc("/products/{id}", getProductByID)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
