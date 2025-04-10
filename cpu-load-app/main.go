package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func healthHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("ok"))
}

type UUIDResponse struct {
	UUID string `json:"uuid"`
	Hash string `json:"hash"`
}

func uuidHandler(w http.ResponseWriter, req *http.Request) {
	newUUID := uuid.New()
	hash := sha256.Sum256(newUUID[:])
	
	response := UUIDResponse{
		UUID: newUUID.String(),
		Hash: hex.EncodeToString(hash[:]),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError) 
		log.Printf("Encoding error: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("GET /uuid", uuidHandler)

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}