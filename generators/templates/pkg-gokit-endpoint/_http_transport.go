package <%- packagename%>

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// MakeHTTPHandler ...
func MakeHTTPHandler(endpoints Endpoints) http.Handler {
	r := mux.NewRouter()

	getHandler := kithttp.NewServer(
		endpoints.GetEndpoint,
		decodeRequest,
		encodeResponse,
	)

	r.Methods("GET").Path("/<%- packagename%>/").Handler(getHandler)

	return r
}

func decodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	params := r.URL.Query()
	id, ok := params["id"]

	if ok {
		if len(id) > 0 && len(id[0]) > 0 {
			return <%- packagename%>Request{ID: id[0]}, nil
		}
		return nil, errors.New("Invalid Argument")
	}

	return <%- packagename%>Request{}, errors.New("Parameter 'id' is not found")
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}
