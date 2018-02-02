package <%- packagename%>

import (
	"context"

	tracing "github.com/ricardo-ch/go-tracing"
)

// <%- packagename%>Tracing struct
type <%- packagename%>Tracing struct {
	next Service
}

// NewTracing ...
func NewTracing(i Service) Service {
	return <%- packagename%>Tracing{
		next: i,
	}
}

// Get <%- packagename%> and tracing info
func (s <%- packagename%>Tracing) Get(ctx context.Context, request <%- packagename%>Request) (response <%- packagename%>Response, err error) {

	span, ctx := tracing.CreateSpan(ctx, "<%- packagename%>.service::Get", &map[string]interface{}{"id": request.ID})
	defer func() {
		if err != nil {
			tracing.SetSpanError(span, err)
		}
		span.Finish()
	}()

	return s.next.Get(ctx, request)
}