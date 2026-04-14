package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/demo", demoHandler)

	log.Println("Hello, World!")
	err := http.ListenAndServe(":8080", nil) // localhost:8080
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}

func demoHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%+v", r)

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := map[string]string{
		"message": "Hello, World!",
		"info":    "Tobyblgo",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
