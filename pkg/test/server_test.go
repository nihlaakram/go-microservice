package test

import (
	"github.com/nihlaakram/go-microservice/pkg/service"
	"log"
	"net/http"
	"net/http/httptest"
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
	dbPass := os.Getenv("TEST_DB_PASS")
	dbName := os.Getenv("TEST_DB_NAME")
	dbUser := os.Getenv("TEST_DB_USER")
	dbHost := os.Getenv("TEST_DB_HOST")
	dbPort := os.Getenv("TEST_DB_PORT")
	server.Init(dbUser, dbPass, dbName, dbHost, dbPort)
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

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	server.Router.ServeHTTP(rec, req)

	return rec
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
