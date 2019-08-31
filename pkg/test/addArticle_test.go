/*
 * Copyright (c) 2019, Nihla Akram. All Rights Reserved.
 */

package test

import (
	"bytes"
	"fmt"
	"github.com/nihlaakram/go-microservice/pkg/util"
	"net/http"
	"net/http/httptest"
	"testing"
)

var article1 = []byte(`{"title": "Hello World",
"content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod ",
"author": "John"}`)

var invalidArticle1 = []byte(`{"content": "Lorem ipsum dolor sit amet, consectetur adipisci",
"author": "John"}`)

var invalidArticle2 = []byte(`{"hello":"world",
content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed ",
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

func TestAddNthArticle(t *testing.T) {
	deleteTableEntries()

	var response *httptest.ResponseRecorder
	for i := 0; i < 10; i++ {
		req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/%v", util.ArticlesResource), bytes.NewBuffer(article1))
		response = executeRequest(req)
	}
	status := http.StatusCreated
	checkResponseCode(t, status, response.Code)

	expectedResponse := `{"status":201,"message":"Success","data":{"id":10}}`
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
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/%v", util.ArticlesResource), bytes.NewBuffer(invalidArticle2))
	response := executeRequest(req)
	status := http.StatusBadRequest
	checkResponseCode(t, status, response.Code)

	expectedResponse := `{"status":400,"message":"Bad Request","data":null}`
	if body := response.Body.String(); body != expectedResponse {
		t.Errorf("Expected %s. Got %s", expectedResponse, body)

	}
}
