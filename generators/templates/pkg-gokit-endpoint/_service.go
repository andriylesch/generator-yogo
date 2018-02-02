package <%- packagename%>

import (
	"context"
)

type Service interface {
	Get(ctx context.Context, request <%- packagename%>Request) (<%- packagename%>Response, error)
}

// Service struct
type service struct {
}

// NewService returns a new instance of the default Service.
func NewService() (Service, error) {
	service := &service{}
	return service, nil
}

func (s *service) Get(ctx context.Context, request <%- packagename%>Request) (<%- packagename%>Response, error) {

	response := <%- packagename%>Response{
		ID: request.ID,
	}

	return response, nil
}
