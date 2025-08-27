package cmd

import (
	"e-commerce/global_router"
	"e-commerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	manager := middleware.NewManager()

	manager.Use(middleware.Logger, middleware.Hudai)

	initRoutes(mux, manager)

	port := ":8080"
	fmt.Println("Starting server on ", port)

	globalRouter := global_router.GlobalRouter(mux)

	err := http.ListenAndServe(port, globalRouter)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
