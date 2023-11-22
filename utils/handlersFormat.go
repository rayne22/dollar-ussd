package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JsonHandler handles json responses
func JsonHandler(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		_, _ = fmt.Fprintf(w, "%s", err.Error())
	}
}

// ErrorHandler Error handles error responses
func ErrorHandler(w http.ResponseWriter, statusCode int, err error, resp map[string]interface{}) {
	if err != nil {
		JsonHandler(w, statusCode, resp)
		return
	}
	JsonHandler(w, http.StatusBadRequest, nil)
}
