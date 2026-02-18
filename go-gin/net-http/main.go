package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	response := map[string]string{"status": "ok"}
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	response := map[string]string{
		"message": "Hello Daniel",
		"user":    "Daniel",
		"status":  "success",
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response) // This is a more concise way to write the JSON response, and it also handles errors internally.
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		fmt.Fprintf(w, "Ok")
	})

	http.HandleFunc("/hello", helloHandler)

	http.HandleFunc("/health", healthCheckHandler)

	log.Println("Starting server on http://localhost:8080")
	// :8080 = localhost:8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		// fatalf = log the error and exit the program
		log.Fatalf("Server failed to start: %v", err)
	}
}
