package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

type Me struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Github string `json:"github"`
}

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// 🔥 CRITICAL FIX: enforce exact match
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, 200, Message{Message: "API is running"})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/health" {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, 200, Message{Message: "healthy"})
}

func meHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/me" {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, 200, Me{
		Name:   "Nafiu Atanda",
		Email:  "atanda0x@gmail.com",
		Github: "https://github.com/0xatanda",
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/me", meHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Server running on port 8080")
	log.Fatal(server.ListenAndServe())
}
