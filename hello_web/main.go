package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.DefaultServeMux

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "web"
		}

		response := make(map[string]interface{})
		response["message"] = fmt.Sprintf("Hello %s!", name)

		log.Printf(`Responding with "%s"`, response["message"])

		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Printf("Failed to write response: %v", err)
		}
	})

	log.Println("Running and listening on port 8080...")
	http.ListenAndServe(":8080", mux)
}
