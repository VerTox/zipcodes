package v1

import (
	"encoding/json"
	"net/http"
)

type ListResponse struct {
	Items interface{} `json:"items"`
	Total int         `json:"total"`
}

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func JsonError(w http.ResponseWriter, status int, err error) {
	j, _ := json.Marshal(&ErrorResponse{
		Error: err.Error(),
		Code:  status,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(j)
}

func JsonSuccessCreated(w http.ResponseWriter, r interface{}) {
	j, _ := json.Marshal(r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func JsonSuccess(w http.ResponseWriter, r interface{}) {
	j, _ := json.Marshal(r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
