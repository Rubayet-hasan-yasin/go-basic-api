package cmd

import (
	"ecommerce/config"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	
	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()
	
	middleware := middleware.NewMiddlewares(cnf)
	
	productHandler := product.NewHandler(middleware, productRepo)
	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(cnf, productHandler, userHandler)

	server.Start()
}
