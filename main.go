package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		Message: "Hello world from json",
		Status:  200,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(resp)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request received: %s %s\n", r.Method, r.URL.Path)
		helloHandler(w, r)
	})

	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("JSON request received: %s %s\n", r.Method, r.URL.Path)
		jsonHandler(w, r)
	})
	fmt.Println("Starting server on :8080..")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
