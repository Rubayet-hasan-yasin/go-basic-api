package rest

import (
	"ecommerce/config"
	"ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Start(cnf config.Config) {

	mux := http.NewServeMux()

	manager := middleware.NewManager()

	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	port := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("Starting server on ", port)

	err := http.ListenAndServe(port, wrappedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
