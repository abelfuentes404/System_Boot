package middleware

import (
    "net/http"
    "github.com/gorilla/handlers"
)

// CORSMiddleware permisivo (SOLO PARA DESARROLLO)
func CORSMiddleware(router http.Handler) http.Handler {
	// Define las opciones de CORS usando gorilla/handlers
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:5173"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	allowCredentials := handlers.AllowCredentials()

	// Devuelve el handler de gorilla/handlers configurado, que envuelve tu enrutador principal
	return handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders, allowCredentials)(router)
}

//para produccion
/*
// CORSMiddleware permite requests desde React (puerto 5173)
func CORSMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Permitir desde localhost:5173 (Vite dev server)
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
        w.Header().Set("Access-Control-Allow-Credentials", "true")

        // Manejar preflight requests
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}*/
