package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	productHandler "ecommerce/rest/handlers/product"
	userHandler "ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
	"ecommerce/user"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbConn, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbConn, "./migrations")
	if err != nil {
		fmt.Println("Failed to apply migrations:", err)
		os.Exit(1)
	}

	//repo
	productRepo := repo.NewProductRepo(dbConn)
	userRepo := repo.NewUserRepo(dbConn)
	
	//domains
	userService := user.NewService(userRepo)

	middleware := middleware.NewMiddlewares(cnf)
	
	productHandler := productHandler.NewHandler(middleware, productRepo)
	userHandler := userHandler.NewHandler(cnf, userService)

	server := rest.NewServer(cnf, productHandler, userHandler)

	server.Start()
}
