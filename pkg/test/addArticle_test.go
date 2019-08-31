package test

import (
	"bytes"
	"fmt"
	"github.com/nihlaakram/go-microservice/pkg/util"
	"net/http"
	"testing"
)

var article1 = []byte(`{"title": "Hello World",
"content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et ",
"author": "John"}`)

var invalidArticle1 = []byte(`{"content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et ",
"author": "John"}`)

var invalidArticle2 = []byte(`{"hello":"world",content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et ",
"author": "John"}`)

func TestAddArticle(t *testing.T) {
	deleteTableEntries()
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/%v", util.ArticlesResource), bytes.NewBuffer(article1))
	response := executeRequest(req)
	status := http.StatusCreated
	checkResponseCode(t, status, response.Code)

	expectedResponse := `{"status":201,"message":"Success","data":{"id":1}}`
	if body := response.Body.String(); body != expectedResponse {
		t.Errorf("Expected %s. Got %s", expectedResponse, body)

	}
}

func TestAddInvalidArticle(t *testing.T) {
	deleteTableEntries()
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/%v", util.ArticlesResource), bytes.NewBuffer(invalidArticle1))
	response := executeRequest(req)
	status := http.StatusBadRequest
	checkResponseCode(t, status, response.Code)

	expectedResponse := `{"status":400,"message":"Bad Request","data":null}`
	if body := response.Body.String(); body != expectedResponse {
		t.Errorf("Expected %s. Got %s", expectedResponse, body)

	}
}

func TestAddInvalidArticle2(t *testing.T) {
	deleteTableEntries()
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/%v", util.ArticlesResource), bytes.NewBuffer(invalidArticle1))
	response := executeRequest(req)
	status := http.StatusBadRequest
	checkResponseCode(t, status, response.Code)

	expectedResponse := `{"status":400,"message":"Bad Request","data":null}`
	if body := response.Body.String(); body != expectedResponse {
		t.Errorf("Expected %s. Got %s", expectedResponse, body)

	}
}
