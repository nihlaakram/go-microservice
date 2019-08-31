package main

import "github.com/nihlaakram/go-microservice/pkg/service"

func main() {

	dbUser := ""
	dbPass := ""
	dbName := ""
	port := 8080

	service := service.Server{}
	service.Init(dbUser, dbPass, dbName, "localhost", "3306")
	service.Start(port)
}
