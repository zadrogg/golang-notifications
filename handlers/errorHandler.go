package handlers

import (
	"encoding/json"
	"net/http"
	"os"
)

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Error(error error) {
	if error != nil {
		_, _ = os.Stderr.WriteString(error.Error())
	}
}

func prepareErrorResponse(code int, message string) *errorResponse {
	return &errorResponse{
		Code:    code,
		Message: message,
	}
}

func sendThrow(w http.ResponseWriter, err error, code int) {
	Error(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_ = json.NewEncoder(w).Encode(prepareErrorResponse(code, err.Error()))
}

func Throw500(w http.ResponseWriter, err error) {
	sendThrow(w, err, http.StatusInternalServerError)
}

func Throw400(w http.ResponseWriter, err error) {
	sendThrow(w, err, http.StatusBadRequest)
}
