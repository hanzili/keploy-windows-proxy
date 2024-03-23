package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Assuming the body is JSON, unmarshal into a map for flexibility
	var requestData map[string]interface{}
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Error parsing request data", http.StatusInternalServerError)
		return
	}

	// Here you can process and store the requestData as needed
	// For example, log the request URL
	log.Printf("Received request for URL: %s", requestData["url"])

	// Respond to `mitmproxy` to indicate receipt
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
