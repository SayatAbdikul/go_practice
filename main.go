package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Define a struct for the JSON response
type Response struct {
	Data string `json:"data"`
}

func main() {
	// Define the HTTP handler function
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Create a Response struct with the desired data
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		response := Response{Data: "sayat"}

		// Convert the struct to JSON
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the Content-Type header to indicate JSON
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response to the HTTP response writer
		w.Write(jsonResponse)
	})

	// Start the HTTP server and listen on localhost:8080
	fmt.Println("Server listening on http://localhost:9090")
	http.ListenAndServe(":9090", nil)
}
