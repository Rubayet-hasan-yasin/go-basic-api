package db

import (
	"ecommerce/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cnf config.DBConfig) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cnf.Host,
		cnf.Port,
		cnf.Username,
		cnf.Password,
		cnf.DbName,
		cnf.SSLMode,
	)
}

func NewConnection(cnf config.DBConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(cnf)
	db, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return nil, err
	}
	fmt.Println("Successfully connected to the database")
	return db, nil
}
