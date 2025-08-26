package main

import (
    "log"
    "net/http"
    
    "time"

    "system_boot/internal/config"
    "system_boot/internal/handlers"
    "system_boot/internal/middleware"
    "system_boot/internal/storage"

    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
)

func main() {
    // Cargar variables de entorno
    if err := godotenv.Load(); err != nil {
        log.Println("No se encontró archivo .env")
    }

    // Inicializar configuración
    cfg := config.LoadConfig()

    // Inicializar almacenamiento
    store := storage.NewStorage()

    // Cargar estado inicial
    if err := store.LoadState(); err != nil {
        log.Printf("Error cargando estado inicial: %v", err)
    }

    // Configurar router
    r := mux.NewRouter()

    // Middleware de logging
    r.Use(middleware.LoggingMiddleware)

    // Rutas públicas
    r.HandleFunc("/api/health", handlers.HealthCheck).Methods("GET")
    r.HandleFunc("/api/setup/status", handlers.GetSetupStatus(store)).Methods("GET")
    r.HandleFunc("/api/setup/login", handlers.HandleSetupLogin(store)).Methods("POST")
    r.HandleFunc("/api/setup/test-connection", handlers.HandleTestConnection).Methods("POST")
    r.HandleFunc("/api/setup/configure", handlers.HandleSetupConfiguration(store)).Methods("POST")
    r.HandleFunc("/api/setup/create-tables", handlers.HandleCreateTables(store)).Methods("POST")
    r.HandleFunc("/api/setup/create-admin", handlers.HandleCreateAdmin(store)).Methods("POST")
    r.HandleFunc("/api/setup/complete", handlers.HandleSetupComplete(store)).Methods("POST")

    // Rutas protegidas (solo en modo producción)
    protected := r.PathPrefix("/api").Subrouter()
    protected.Use(middleware.JWTAuthMiddleware)
    protected.HandleFunc("/dashboard", handlers.HandleDashboard).Methods("GET")

    // Servidor HTTP
    srv := &http.Server{
        Handler:      r,
        Addr:         ":" + cfg.Port,
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    log.Printf("Servidor iniciado en puerto %s", cfg.Port)
    log.Fatal(srv.ListenAndServe())
}