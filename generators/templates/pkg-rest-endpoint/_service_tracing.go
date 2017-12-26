package <%- packagename%>

import (
	"context"

	tracing "github.com/ricardo-ch/go-tracing"
)

type <%- packagename%>Tracing struct {
	next Service
}

// NewTracing ...
func NewTracing(s Service) Service {
	return <%- packagename%>Tracing{
		next: s,
	}
}

// Get ...
func (s <%- packagename%>Tracing) Get(ctx context.Context, obj <%- packagename%>Request) (response <%- packagename%>Response, err error) {
	span, ctx := tracing.CreateSpan(ctx, "<%- packagename%>.service::Get", &map[string]interface{}{"id": obj.ID})
	defer func() {
		if err != nil {
			tracing.SetSpanError(span, err)
		}
		span.Finish()
	}()

	return s.next.Get(ctx, obj)
}
