package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbConn, err := db.NewConnection()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		os.Exit(1)
	}
	
	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo(dbConn)
	
	middleware := middleware.NewMiddlewares(cnf)
	
	productHandler := product.NewHandler(middleware, productRepo)
	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(cnf, productHandler, userHandler)

	server.Start()
}
