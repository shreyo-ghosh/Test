package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	Time    string `json:"time"`
	Message string `json:"message"`
	Version string `json:"version"`
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Time:    time.Now().Format(time.RFC3339),
		Message: "Hello from Cloud Function!",
		Version: "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
