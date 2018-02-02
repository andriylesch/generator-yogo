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

	options := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError()),
	}

	getHandler := kithttp.NewServer(
		endpoints.GetEndpoint,
		decodeRequest,
		encodeResponse,
		options...,
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

func encodeError() kithttp.ErrorEncoder {
	return func(ctx context.Context, err error, w http.ResponseWriter) {

		// create extra fields for response
		outputBody := map[string]interface{}{}

		// you can write your logic for handle errors

		outputBody["error"] = err.Error()
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(outputBody)
	}
}
