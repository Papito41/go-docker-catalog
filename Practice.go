package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	catalog := []Product{
		{Name: "Leather Shoes", Price: 15000.00},
		{Name: "School Bag", Price: 8500.00},
	}

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(catalog)
	})

	http.ListenAndServe(":8080", nil)
}
