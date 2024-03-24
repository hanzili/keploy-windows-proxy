package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request")
	fmt.Println(r)
	// body, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	http.Error(w, "Error reading request body", http.StatusInternalServerError)
	// 	return
	// }
	// defer r.Body.Close()

	// var requestData map[string]interface{}
	// if err := json.Unmarshal(body, &requestData); err != nil {
	// 	fmt.Println(w, "Error parsing request data", http.StatusInternalServerError)
	// 	return
	// }

	// log.Printf("Received request for URL: %s", requestData["url"])

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
