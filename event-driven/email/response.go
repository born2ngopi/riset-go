package main

import (
	"encoding/json"
	"net/http"
)

type ResponseErrorDTO struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func ResponseNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(ResponseErrorDTO{
		Message: "not found",
		Code:    http.StatusNotFound,
	})
}

func ResponseServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(ResponseErrorDTO{
		Message: "internal server error",
		Code:    http.StatusInternalServerError,
	})
}

func ResponseBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(ResponseErrorDTO{
		Message: "bad request",
		Code:    http.StatusBadRequest,
	})
}
