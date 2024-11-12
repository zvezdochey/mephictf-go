package main

import (
	"fmt"
	"log"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprint(w, "pong!"); err != nil {
		log.Printf("Failed to write response: %s", err.Error())
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", PingHandler)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
	}
}
