/*
 * Copyright (c) 2019, Nihla Akram. All Rights Reserved.
 */

 package test

import (
	"fmt"
	"github.com/nihlaakram/go-microservice/pkg/util"
	"net/http"
	"testing"
)

func TestEmptyTable(t *testing.T) {
	deleteTableEntries()

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/%v", util.ArticlesResource), nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	//expectedResponse, err := json.Marshal(&model.Response{http.StatusBadRequest, util.SuccessMsg, []model.Article{}})
	expectedResponse := `{"status":200,"message":"Success","data":[]}`
	if body := response.Body.String(); body != expectedResponse {
		t.Errorf("Expected %s. Got %s", expectedResponse, body)
	}
}
