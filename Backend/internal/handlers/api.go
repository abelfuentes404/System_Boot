package handlers

import (
    "encoding/json"
    "log"
    "net/http"

    "system_boot/internal/storage"
)

// HandlerFunc con storage
type HandlerFuncWithStore func(http.ResponseWriter, *http.Request, *storage.Storage)

// MakeHandler crea un handler regular desde uno con storage
func MakeHandler(hf HandlerFuncWithStore, store *storage.Storage) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        hf(w, r, store)
    }
}

// HealthCheck endpoint
func HealthCheck(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok", "message": "Server is running"})
}

// respondWithJSON helper function
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    if err := json.NewEncoder(w).Encode(payload); err != nil {
        log.Printf("Error encoding JSON: %v", err)
    }
}

// respondWithError helper function
func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}