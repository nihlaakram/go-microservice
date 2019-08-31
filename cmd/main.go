package main

import (
	"github.com/nihlaakram/go-microservice/pkg/service"
	"os"
)

func main() {

	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	serverPort := 8080

	service := service.Server{}
	service.Init(dbUser, dbPass, dbName, dbHost, dbPort)
	service.Start(serverPort)
}
