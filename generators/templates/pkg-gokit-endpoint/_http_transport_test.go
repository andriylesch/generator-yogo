package <%- packagename%>

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MakeHTTPHandler(t *testing.T) {
	h := MakeHTTPHandler(Endpoints{})
	assert.NotNil(t, h)
}

func Test_decodeRequest_Success(t *testing.T) {
	// Arrange
	expectedResult := testObjRequest

	// Act
	httpRequest, _ := http.NewRequest("GET", fmt.Sprintf("/<%- packagename%>/?id=%s", testID), nil)
	req, err := decodeRequest(context.Background(), httpRequest)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, req)
}

func Test_decodeRequest_With_Empty_ID_Fail(t *testing.T) {
	// Arrange
	expectedError := errors.New("Invalid Argument")

	// Act
	httpRequest, err := http.NewRequest("GET", "/<%- packagename%>?id=", nil)
	_, err = decodeRequest(context.Background(), httpRequest)

	//Assert
	assert.NotNil(t, err)
	assert.EqualError(t, expectedError, err.Error())
}

func Test_decodeRequest_Without_ID_Fail(t *testing.T) {
	// Arrange
	expectedError := errors.New("Parameter 'id' is not found")

	// Act
	httpRequest, err := http.NewRequest("GET", "/<%- packagename%>", nil)
	_, err = decodeRequest(context.Background(), httpRequest)

	//Assert
	assert.NotNil(t, err)
	assert.EqualError(t, expectedError, err.Error())
}
