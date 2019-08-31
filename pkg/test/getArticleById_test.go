package test

import (
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
