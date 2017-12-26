package <%- packagename%>

import (
	"context"
)

type Service interface {
	Get(ctx context.Context, obj <%- packagename%>Request) (<%- packagename%>Response, error)
}

type Repository interface {
	Get(ID string) (<%- packagename%>Response, error)
}
