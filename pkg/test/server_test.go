package test

import (
	"github.com/nihlaakram/go-microservice/pkg/service"
	"log"
	"os"
	"testing"
)

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS articles (
		id INT NOT NULL AUTO_INCREMENT,
		title VARCHAR(45) NULL,
		content TEXT NULL,
		author VARCHAR(45) NULL,
	PRIMARY KEY (id));`

var server service.Server

func TestMain(m *testing.M) {
	server = service.Server{}
	dbPass := ""
	dbName := ""
	//port := 8080
	dbUser := "root"

	server.Init(dbUser, dbPass, dbName, "localhost", "3306")
	checkIfTableExists()
	code := m.Run()
	deleteTableEntries()
	os.Exit(code)
}

func checkIfTableExists() {
	if _, err := server.DBCon.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func deleteTableEntries() {
	server.DBCon.Exec("DELETE FROM articles")
	server.DBCon.Exec("ALTER TABLE articles AUTO_INCREMENT = 1")
}
