package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int     `json:"_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"imageUrl"`
}

var products []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(products)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/products", getProducts)

	port := ":8080"
	fmt.Println("Starting server on ", port)

	err := http.ListenAndServe(port, mux)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func init() {
	p1 := Product{
		ID:          1,
		Title:       "Mango",
		Description: "A delicious tropical fruit",
		Price:       100,
		ImageUrl:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRf0J-W_xQ8nJ2T7SeBHdkUc68NZIE0Zb4woQ&s",
	}

	p2 := Product{
		ID:          2,
		Title:       "Apple",
		Price:       50,
		Description: "A sweet red fruit",
		ImageUrl:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRf0J-W_xQ8nJ2T7SeBHdkUc68NZIE0Zb4woQ&s",
	}

	p3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "A long yellow fruit",
		Price:       30,
		ImageUrl:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQG7ElBNPs-HbYJJOMHRu7lEmphTn8-52FYKw&s",
	}

	p4 := Product{
		ID:          4,
		Title:       "Pineapple",
		Description: "A tropical fruit with a spiky exterior",
		Price:       150,
		ImageUrl:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRf0J-W_xQ8nJ2T7SeBHdkUc68NZIE0Zb4woQ&s",
	}

	p5 := Product{
		ID:          5,
		Title:       "Grapes",
		Description: "A bunch of small round fruits",
		Price:       200,
		ImageUrl:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRf0J-W_xQ8nJ2T7SeBHdkUc68NZIE0Zb4woQ&s",
	}

	products = append(products, p1, p2, p3, p4, p5)
}
