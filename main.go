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
	// allowCors(w)

	// GET(w, r)

	sendData(w, products, http.StatusOK)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	// allowCors(w)

	// Options(w, r)

	// POST(w, r)

	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "plz give me a valid json", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(products) + 1

	products = append(products, newProduct)

	sendData(w, newProduct, http.StatusCreated)
}

func GET(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
func POST(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func Options(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(200)
	}
}

func allowCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

}

func sendData(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", corsMiddleware(http.HandlerFunc(getProducts)))
	mux.Handle("/create-product", corsMiddleware(http.HandlerFunc(createProduct)))

	port := ":8080"
	fmt.Println("Starting server on ", port)

	globalRouter := globalRouter(mux)

	err := http.ListenAndServe(port, globalRouter)
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

func corsMiddleware(next http.Handler) http.Handler {
	handleCors := func(w http.ResponseWriter, r *http.Request) {
		allowCors(w)
		next.ServeHTTP(w, r)
	}

	handler := http.HandlerFunc(handleCors)

	return handler

}

func globalRouter(mux *http.ServeMux) http.Handler {
	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			allowCors(w)
			w.WriteHeader(http.StatusOK)
		} else {
			mux.ServeHTTP(w, r)
		}
	}

	return http.HandlerFunc(handleAllReq)
}
