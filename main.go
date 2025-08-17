package main

import "net/http"
import "fmt"

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", helloWorld)

	fmt.Println("Starting server on :3000")

	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
