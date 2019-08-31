package test

import (
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
