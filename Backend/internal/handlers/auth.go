package handlers

import (
    "encoding/json"
    "net/http"

   
)

// HealthCheck endpoint de verificación de estado
func HealthCheck(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "status":  "ok",
        "message": "Server is running",
    })
}

// HandleDashboard maneja las solicitudes al dashboard
func HandleDashboard(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Welcome to the dashboard",
    })
}