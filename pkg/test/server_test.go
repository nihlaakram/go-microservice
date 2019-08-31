/*
 * Copyright (c) 2019, Nihla Akram. All Rights Reserved.
 */

package test

import (
	"github.com/nihlaakram/go-microservice/pkg/service"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const deleteTableData = "DELETE FROM articles"
const resetArticleId = "ALTER TABLE articles AUTO_INCREMENT = 1"

var server service.Server

func TestMain(m *testing.M) {
	server = service.Server{}
	dbPass := os.Getenv("TEST_DB_PASS")
	dbName := os.Getenv("TEST_DB_NAME")
	dbUser := os.Getenv("TEST_DB_USER")
	dbHost := os.Getenv("TEST_DB_HOST")
	dbPort := os.Getenv("TEST_DB_PORT")

	server.Init(dbUser, dbPass, dbName, dbHost, dbPort)
	code := m.Run()
	deleteTableEntries()
	os.Exit(code)
}

func deleteTableEntries() {
	server.DBCon.Exec(deleteTableData)
	server.DBCon.Exec(resetArticleId)
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
