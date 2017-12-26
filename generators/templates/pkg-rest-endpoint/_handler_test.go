package <%- packagename%>

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var testResult = <%- packagename%>Response{
	ID: "1234567",
}

type testService struct {
	Service
}

func (m testService) Get(obj <%- packagename%>Request) (<%- packagename%>Response, error) {
	return testResult, nil
}

var url = "/<%- packagename%>/1234567"

func TestGet_Handler(t *testing.T) {
	//arrange
	request, _ := http.NewRequest("GET", url, nil)
	response := httptest.NewRecorder()
	json, _ := json.Marshal(testResult)
	expectedJSON := string(json)

	//act
	r := mux.NewRouter()
	r.HandleFunc(url, NewHandler(testService{}).Get)
	r.ServeHTTP(response, request)

	//assert
	assert.Equal(t, "application/json; charset=UTF-8", response.Header().Get("Content-Type"))
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, strings.TrimSpace(expectedJSON), strings.TrimSpace(response.Body.String()))

}
