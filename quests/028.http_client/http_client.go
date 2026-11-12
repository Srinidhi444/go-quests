package http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func newHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 5 * time.Second,
	}
}

// FetchProducts retrieves all products.
func FetchProducts(apiURL string) ([]Product, error) {
	client := newHTTPClient()

	resp, err := client.Get(apiURL + "/products")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var products []Product
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil, err
	}

	return products, nil
}

// CreateProduct sends a new product to the server.
func CreateProduct(apiURL string, p Product) error {
	client := newHTTPClient()

	var body bytes.Buffer

	if err := json.NewEncoder(&body).Encode(p); err != nil {
		return err
	}

	resp, err := client.Post(
		apiURL+"/products",
		"application/json",
		&body,
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated &&
		resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	return nil
}

// FetchProduct retrieves a product by ID.
func FetchProduct(apiURL string, id string) (Product, bool, error) {
	client := newHTTPClient()

	resp, err := client.Get(apiURL + "/products/" + id)
	if err != nil {
		return Product{}, false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return Product{}, false, nil
	}

	if resp.StatusCode != http.StatusOK {
		return Product{}, false, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var product Product
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		return Product{}, false, err
	}

	return product, true, nil
}

// FetchMultipleProducts retrieves products concurrently.
func FetchMultipleProducts(apiURL string, ids []string) (map[string]Product, error) {
	results := make(map[string]Product)

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, id := range ids {
		wg.Add(1)

		go func(id string) {
			defer wg.Done()

			product, found, err := FetchProduct(apiURL, id)

			// Skip missing products and failed requests.
			if err != nil || !found {
				return
			}

			mu.Lock()
			results[id] = product
			mu.Unlock()
		}(id)
	}

	wg.Wait()

	return results, nil
}
