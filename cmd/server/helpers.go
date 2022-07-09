package server

import (
	"encoding/json"
	"net/http"
)

func WriteErr(w http.ResponseWriter, statusCode int, reason string) {
	w.WriteHeader(statusCode)

	err := Map{
		"error": reason,
	}

	ToJSON(w, err)
} 

func ToJSON(w http.ResponseWriter, data any) {
	json.NewEncoder(w).Encode(data)
}