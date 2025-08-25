package cmd

import (
	"e-commerce/global_router"
	"e-commerce/handlers"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductById))

	port := ":8080"
	fmt.Println("Starting server on ", port)

	globalRouter := global_router.GlobalRouter(mux)

	err := http.ListenAndServe(port, globalRouter)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
