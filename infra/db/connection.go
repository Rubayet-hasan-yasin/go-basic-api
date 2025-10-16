package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString() string {
	return "host=localhost port=5432 user=nestuser password=nestpass dbname=go_e_com sslmode=disable"
}

func NewConnection() (*sqlx.DB, error) {
	dbSource := GetConnectionString()
	db, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return nil, err
	}
	fmt.Println("Successfully connected to the database")
	return db, nil
}