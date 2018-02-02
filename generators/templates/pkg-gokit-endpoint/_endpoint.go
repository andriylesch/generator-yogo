package <%- packagename%>

import "context"
import "github.com/go-kit/kit/endpoint"

// Endpoints represents all endpoints
type Endpoints struct {
	GetEndpoint endpoint.Endpoint
}

func MakeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(<%- packagename%>Request)
		return s.Get(ctx, req)
	}
}
