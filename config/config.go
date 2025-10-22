package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Config

type DBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
	SSLMode  string
}

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
	DB           DBConfig
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load env: ", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is Requird")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service Name is Requird")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("Http Port is Requird")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("Prot must be a number")
		os.Exit(1)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("JWT SECRET KEY is Requird")
		os.Exit(1)
	}

	db_host := os.Getenv("DB_HOST")
	if db_host == "" {
		fmt.Println("DB_HOST is Requird")
		os.Exit(1)
	}

	db_port := os.Getenv("DB_PORT")
	if db_port == "" {
		fmt.Println("DB_PORT is Requird")
		os.Exit(1)
	}
	portInt, err := strconv.Atoi(db_port)
	if err != nil {
		fmt.Println("DB_PORT must be a number")
		os.Exit(1)
	}
	db_username := os.Getenv("DB_USERNAME")
	if db_username == "" {
		fmt.Println("DB_USERNAME is Requird")
		os.Exit(1)
	}
	db_password := os.Getenv("DB_PASSWORD")
	if db_password == "" {
		fmt.Println("DB_PASSWORD is Requird")
		os.Exit(1)
	}

	db_name := os.Getenv("DB_NAME")
	if db_name == "" {
		fmt.Println("DB_NAME is Requird")
		os.Exit(1)
	}

	db_sslmode := os.Getenv("DB_SSLMODE")
	if db_sslmode == "" {
		fmt.Println("DB_SSLMODE is Requird")
		os.Exit(1)
	}


	dbConfig := DBConfig{
		Host:     db_host,
		Port:     portInt,
		Username: db_username,
		Password: db_password,
		DbName:   db_name,
		SSLMode:  db_sslmode,
	}

	configurations = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     int(port),
		JwtSecretKey: jwtSecretKey,
		DB:           dbConfig,
	}
}

func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}
	return configurations
}
