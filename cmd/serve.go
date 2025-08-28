package cmd

import (
	"e-commerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	manager := middleware.NewManager()

	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	port := ":8080"
	fmt.Println("Starting server on ", port)

	// globalRouter := global_router.GlobalRouter(mux)

	err := http.ListenAndServe(port, wrappedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
