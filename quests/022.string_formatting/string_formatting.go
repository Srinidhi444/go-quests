package string_formatting

import "fmt"

type Product struct {
	name      string
	quantity  int
	price     float64
	available bool
}

func StringFormating() {
	p := Product{
		name:      "Laptop",
		quantity:  42,
		price:     899.99,
		available: true,
	}

	fmt.Printf("Product: %+v\n", p)
	fmt.Printf("Details: %+v\n", p)
	fmt.Printf("Go-syntax: %#v\n", p)
	fmt.Printf("Type: %T\n", p)
	fmt.Printf("Available: %t\n", p.available)
	fmt.Printf("Quantity: %d\n", p.quantity)
	fmt.Printf("Quantity in binary: %b\n", p.quantity)
	fmt.Printf("Quantity in hex: %x\n", p.quantity)
	fmt.Printf("Price: $%.2f\n", p.price)
	fmt.Printf("Formatted: |%-10s|%5d|%7.2f|\n", p.name, p.quantity, p.price)
}
