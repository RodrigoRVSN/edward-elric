package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		return "6969"
	}

	return port
}

func handleRequest(writter http.ResponseWriter, r *http.Request) {
	writter.Header().Set("Content-Type", "application/json")

	response := make(map[string]string)
	response["message"] = "Ze da mangaaaaa papai"
	jsonResponse, error := json.Marshal(response)

	if error != nil {
		fmt.Fprintf(writter, "Ih rapaz")
	}

	writter.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/health-check", handleRequest)

	PORT := getPort()

	fmt.Println("🚀 Server running on port: " + PORT)
	http.ListenAndServe(":"+PORT, nil)
}
