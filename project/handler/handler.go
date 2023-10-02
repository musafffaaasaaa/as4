package handler

import (
	"context"
	"encoding/json"
	"net/http"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	// Add other fields as needed
}

func NewHandler() {
	// Здесь вы можете выполнить все необходимые настройки или конфигурацию.
	// Например, инициализацию подключений к базе данных, настройку промежуточного ПО и т. д.
}

// CreateProductHandler handles the creation of a new product.
func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product

	// Parse JSON data from the request into the product struct
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get a MongoDB client
	client, err := db.ConnectToMongoDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.TODO())

	// Insert the product into the database
	err = db.InsertProduct(client, product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetProductHandler retrieves a product by ID.
func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement logic to retrieve a product by ID from the database.
	// Replace the following example code with your database query logic.
	productID := 1
	product := findProductByID(productID)

	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(product); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func findProductByID(id int) *Product {
	// TODO: Implement logic to find a product by ID from the database.
	// Replace the following example code with your database query logic.
	products := []Product{
		{ID: 1, Name: "Product 1", Price: 19.99},
		{ID: 2, Name: "Product 2", Price: 29.99},
	}

	for _, p := range products {
		if p.ID == id {
			return &p
		}
	}
	return nil
}
