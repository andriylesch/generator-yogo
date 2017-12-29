package <%- packagename%>

import "context"

//Service...
type service struct {
	Repository Repository
}

//NewService ...
func NewService(r Repository) Service {
	return service{Repository: r}
}

//Get ...
func (s service) Get(ctx context.Context, obj <%- packagename%>Request) (<%- packagename%>Response, error) {
	return s.Repository.Get(obj.ID)
}
