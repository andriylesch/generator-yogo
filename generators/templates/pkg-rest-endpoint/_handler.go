package <%- packagename%>

import (
	"encoding/json"
	"net/http"

	logger "github.com/ricardo-ch/go-logger"
)

//Handler ...
type Handler struct {
	Service Service
}

//NewHandler ...
func NewHandler(s Service) Handler {
	return Handler{Service: s}
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	params := r.URL.Query()
	id, ok := params["id"]

	var result interface{}
	var err error

	if ok {
		if len(id) > 0 && len(id[0]) > 0 {
			result, err = h.Service.Get(r.Context(), <%- packagename%>Request{ID: id[0]})
			if err != nil {
				logger.Error(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
