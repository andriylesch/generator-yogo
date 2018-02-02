package <%- packagename%>

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	testID = "test-id"
)

var testError = errors.New("test error")

var testObjRequest = <%- packagename%>Request{
	ID: testID,
}

var testObjResponse = <%- packagename%>Response{
	ID: testID,
}

type serviceMock struct {
	mock.Mock
}

func (s serviceMock) Get(ctx context.Context, request <%- packagename%>Request) (<%- packagename%>Response, error) {
	args := s.Called(request)
	return args.Get(0).(<%- packagename%>Response), args.Error(1)
}

func TestGet_Success(t *testing.T) {
	// Arrange
	service := &serviceMock{}
	serviceTracing := NewTracing(service)
	service.On("Get", testObjRequest).Return(testObjResponse, nil)

	//Act
	result, err := serviceTracing.Get(context.Background(), testObjRequest)

	//Assert
	assert.NoError(t, err)
	assert.NotNil(t, result, "result should not be nil")
	assert.Equal(t, result, testObjResponse)
}

func TestGet_Fail(t *testing.T) {
	// Arrange
	service := &serviceMock{}
	serviceTracing := NewTracing(service)
	service.On("Get", testObjRequest).Return(testObjResponse, testError)

	//Act
	result, err := serviceTracing.Get(context.Background(), testObjRequest)

	//Assert
	assert.Error(t, err)
	assert.NotNil(t, result, "result should not be nil")
	assert.EqualError(t, err, testError.Error())
}
