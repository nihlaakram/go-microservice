/*
 * Copyright (c) 2019, Nihla Akram. All Rights Reserved.
 */

package test

import (
	"bytes"
	"fmt"
	"github.com/nihlaakram/go-microservice/pkg/util"
	"net/http"
	"testing"
)

func TestGetNonExistentArticle(t *testing.T) {
	deleteTableEntries()

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/%v/%d", util.ArticlesResource, 1), nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	expectedResponse := `{"status":404,"message":"No such article found","data":null}`
	if body := response.Body.String(); body != string(expectedResponse) {
		t.Errorf("Expected %s. Got %s", string(expectedResponse), body)
	}
}

func TestGetExistentArticle(t *testing.T) {
	deleteTableEntries()

	article1 := []byte(`{"title":"title1","content":"content1","author":"author1"}`)
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/%v", util.ArticlesResource), bytes.NewBuffer(article1))
	response := executeRequest(req)

	req, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/%v/%d", util.ArticlesResource, 1), nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	expectedResponse := `{"status":200,"message":"Success","data":{"id":1,"title":"title1","content":"content1","author":"author1"}}`
	if body := response.Body.String(); body != string(expectedResponse) {
		t.Errorf("Expected %s. Got %s", string(expectedResponse), body)
	}
}

func TestGetInvalidArticle(t *testing.T) {
	deleteTableEntries()

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/%v/abc", util.ArticlesResource), nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)

	expectedResponse := `{"status":400,"message":"Invalid article Id","data":null}`
	if body := response.Body.String(); body != string(expectedResponse) {
		t.Errorf("Expected %s. Got %s", string(expectedResponse), body)
	}
}
