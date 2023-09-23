package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/health-check", handleRequest)

	http.ListenAndServe(":6969", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := make(map[string]string)
	response["message"] = "Mim de papai"
	jsonResponse, error := json.Marshal(response)

	if error != nil {
		fmt.Fprintf(w, "Ih rapaz")
	}

	w.Write(jsonResponse)
}
