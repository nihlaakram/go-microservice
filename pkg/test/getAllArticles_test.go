package test

import (
	"bytes"
	"fmt"
	"github.com/nihlaakram/go-microservice/pkg/util"
	"net/http"
	"testing"
)

func TestGetNonExistentAllArticles(t *testing.T) {
	deleteTableEntries()

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/%v", util.ArticlesResource), nil)
	response := executeRequest(req)
	status := http.StatusOK
	checkResponseCode(t, status, response.Code)
	expectedResponse := `{"status":200,"message":"Success","data":[]}`
	if body := response.Body.String(); body != expectedResponse {
		t.Errorf("Expected %s. Got %s", expectedResponse, body)
	}
}

func TestTableWithOneEntry(t *testing.T) {
	deleteTableEntries()
	// add data
	article1 := []byte(`{"title":"title1","content":"content1","author":"author1"}`)
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/%v", util.ArticlesResource), bytes.NewBuffer(article1))
	response := executeRequest(req)

	req, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/%v", util.ArticlesResource), nil)
	response = executeRequest(req)
	status := http.StatusOK
	checkResponseCode(t, status, response.Code)
	//expectedResponse, err := json.Marshal(&model.Response{http.StatusBadRequest, util.SuccessMsg, []model.Article{}})
	expectedResponse := `{"status":200,"message":"Success","data":[{"id":1,"title":"title1","content":"content1","author":"author1"}]}`
	if body := response.Body.String(); body != expectedResponse {
		t.Errorf("Expected %s. Got %s", expectedResponse, body)
	}
}

func TestTableWithMultipleEntry(t *testing.T) {
	deleteTableEntries()
	// add data
	article1 := []byte(`{"title":"title1","content":"content1","author":"author1"}`)
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/%v", util.ArticlesResource), bytes.NewBuffer(article1))
	response := executeRequest(req)
	req, _ = http.NewRequest(http.MethodPost, fmt.Sprintf("/%v", util.ArticlesResource), bytes.NewBuffer(article1))
	response = executeRequest(req)

	req, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/%v", util.ArticlesResource), nil)
	response = executeRequest(req)
	status := http.StatusOK
	checkResponseCode(t, status, response.Code)
	//expectedResponse, err := json.Marshal(&model.Response{http.StatusBadRequest, util.SuccessMsg, []model.Article{}})
	expectedResponse := `{"status":200,"message":"Success","data":[{"id":1,"title":"title1","content":"content1","author":"author1"},{"id":2,"title":"title1","content":"content1","author":"author1"}]}`
	if body := response.Body.String(); body != expectedResponse {
		t.Errorf("Expected %s. Got %s", expectedResponse, body)
	}
}
