package main

import (
    "log"
    "net/http"
    "time"

    "github.com/gorilla/mux"
    "github.com/joho/godotenv"

    "system_boot/internal/config"
    "system_boot/internal/handlers"
    "system_boot/internal/middleware"
    "system_boot/internal/storage"
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

    // Middleware
    r.Use(middleware.CORSMiddleware)
    r.Use(middleware.LoggingMiddleware)
    

    // API routes
    api := r.PathPrefix("/api").Subrouter()
    
    // Rutas públicas
    api.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
    api.HandleFunc("/setup/status", handlers.MakeHandler(handlers.GetSetupStatus, store)).Methods("GET")
    api.HandleFunc("/setup/login", handlers.MakeHandler(handlers.HandleSetupLogin, store)).Methods("POST")
    api.HandleFunc("/setup/test-connection", handlers.MakeHandler(handlers.HandleTestConnection, store)).Methods("POST")
    api.HandleFunc("/setup/configure", handlers.MakeHandler(handlers.HandleSetupConfiguration, store)).Methods("POST")
    api.HandleFunc("/setup/create-tables", handlers.MakeHandler(handlers.HandleCreateTables, store)).Methods("POST")
    api.HandleFunc("/setup/create-admin", handlers.MakeHandler(handlers.HandleCreateAdmin, store)).Methods("POST")
    api.HandleFunc("/setup/complete", handlers.MakeHandler(handlers.HandleSetupComplete, store)).Methods("POST")

    // Rutas de autenticación
    api.HandleFunc("/auth/login", handlers.MakeHandler(handlers.HandleAuthLogin, store)).Methods("POST")
    api.HandleFunc("/auth/logout", handlers.HandleAuthLogout).Methods("POST")
    api.HandleFunc("/auth/verify", handlers.MakeHandler(handlers.HandleVerifyToken, store)).Methods("GET")

    // Rutas protegidas
    protected := api.PathPrefix("/protected").Subrouter()
    protected.Use(middleware.JWTAuthMiddleware)
    protected.HandleFunc("/dashboard", handlers.MakeHandler(handlers.HandleDashboardData, store)).Methods("GET")

    // Servir archivos estáticos para React (en producción)
    if cfg.Env == "production" {
        r.PathPrefix("/").Handler(http.FileServer(http.Dir("../frontend/dist/")))
    }

    // Servidor HTTP
    srv := &http.Server{
        Handler:      middleware.CORSMiddleware(r),
        Addr:         ":" + cfg.Port,
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    log.Printf("Servidor iniciado en puerto %s", cfg.Port)
    log.Fatal(srv.ListenAndServe())
}