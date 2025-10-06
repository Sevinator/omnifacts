package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func WriteSuccess(w http.ResponseWriter, status int, data interface{}) {
	WriteJSON(w, status, Response{
		Success: true,
		Data:    data,
	})
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, Response{
		Success: false,
		Error:   message,
	})
}
