package http_utils

import (
	"encoding/json"
	"github.com/ZerepL/bookstore_utils/internal_errors"
	"net/http"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func RespondError(w http.ResponseWriter, err internal_errors.RestErr) {
	RespondJson(w, err.Status(), err)
}
