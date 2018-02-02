package <%- packagename%>

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeEndpoint(t *testing.T) {
	// Arrange
	service := &serviceMock{}
	service.On("Get", testObjRequest).Return(testObjResponse, nil)

	endpoint := MakeEndpoint(service)

	//Act
	wl, err := endpoint(nil, testObjRequest)

	//Assert
	assert.NoError(t, err)
	result, ok := wl.(<%- packagename%>Response)
	assert.Equal(t, true, ok)
	assert.Equal(t, testObjResponse, result)
}
