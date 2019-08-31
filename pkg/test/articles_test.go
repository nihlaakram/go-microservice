package test

import (
	"encoding/json"
	"fmt"
	"github.com/nihlaakram/go-microservice/pkg/model"
	"github.com/nihlaakram/go-microservice/pkg/util"
	"net/http"
	"testing"
)

func TestEmptyTable(t *testing.T) {
	deleteTableEntries()

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/%v", util.ArticlesResource), nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	expectedResponse, err := json.Marshal(&model.Response{http.StatusOK, util.SuccessMsg, []model.Article{}})
	if body := response.Body.String(); body != string(expectedResponse) {
		t.Errorf("Expected %s. Got %s", expectedResponse, body)
	} else if err != nil {
		t.Error(err)
	}
}
