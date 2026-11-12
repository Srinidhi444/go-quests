package http_server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func StartCatalogServer(port int) error {
	// In-memory store
	mp := make(map[string]Product)
	var mu sync.RWMutex

	mux := http.NewServeMux()

	// GET /products
	mux.HandleFunc("GET /products", func(w http.ResponseWriter, r *http.Request) {
		mu.RLock()

		products := make([]Product, 0, len(mp))
		for _, product := range mp {
			products = append(products, product)
		}

		mu.RUnlock()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	})

	// GET /products/{id}
	mux.HandleFunc("GET /products/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		mu.RLock()
		product, exists := mp[id]
		mu.RUnlock()

		if !exists {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
	})

	// POST /products
	mux.HandleFunc("POST /products", func(w http.ResponseWriter, r *http.Request) {
		var product Product

		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		mu.Lock()
		mp[product.ID] = product
		mu.Unlock()

		w.WriteHeader(http.StatusCreated)
	})

	// Start server
	return http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
