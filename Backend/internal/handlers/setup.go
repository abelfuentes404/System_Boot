package handlers

import (
    "encoding/json"
    "net/http"

    "system_boot/internal/storage"
)

// GetSetupStatus devuelve el estado actual del setup
func GetSetupStatus(store *storage.Storage) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        status := map[string]interface{}{
            "setupComplete": store.State.SetupComplete,
            "hasDbConfig":   store.HasDBConfig(),
        }
        
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(status)
    }
}

// HandleSetupLogin maneja el login durante el setup
func HandleSetupLogin(store *storage.Storage) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var credentials struct {
            Username string `json:"username"`
            Password string `json:"password"`
        }
        
        if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }
        
        // Verificar credenciales locales
        if !store.ValidateLocalCredentials(credentials.Username, credentials.Password) {
            http.Error(w, "Invalid credentials", http.StatusUnauthorized)
            return
        }
        
        // Generar token JWT temporal para el setup
        token, err := GenerateSetupToken(credentials.Username)
        if err != nil {
            http.Error(w, "Error generating token", http.StatusInternalServerError)
            return
        }
        
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "token": token,
            "message": "Login successful",
        })
    }
}

// HandleTestConnection prueba la conexión a PostgreSQL
func HandleTestConnection(store *storage.Storage) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var config struct {
            Host     string `json:"host"`
            Port     int    `json:"port"`
            User     string `json:"user"`
            Password string `json:"password"`
            DBName   string `json:"dbname"`
        }
        
        if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }
        
        // Probar conexión
        if err := store.TestPostgresConnection(config.Host, config.Port, config.User, config.Password, config.DBName); err != nil {
            http.Error(w, "Connection failed: "+err.Error(), http.StatusBadRequest)
            return
        }
        
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "message": "Connection successful",
        })
    }
}

// HandleSetupConfiguration guarda la configuración de la base de datos
func HandleSetupConfiguration(store *storage.Storage) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var config struct {
            Host     string `json:"host"`
            Port     int    `json:"port"`
            User     string `json:"user"`
            Password string `json:"password"`
            DBName   string `json:"dbname"`
        }
        
        if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }
        
        // Guardar configuración
        if err := store.SaveDBConfiguration(config.Host, config.Port, config.User, config.Password, config.DBName); err != nil {
            http.Error(w, "Error saving configuration: "+err.Error(), http.StatusInternalServerError)
            return
        }
        
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "message": "Configuration saved successfully",
        })
    }
}

// HandleCreateTables crea las tablas en la base de datos
func HandleCreateTables(store *storage.Storage) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if err := store.CreateTables(); err != nil {
            http.Error(w, "Error creating tables: "+err.Error(), http.StatusInternalServerError)
            return
        }
        
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "message": "Tables created successfully",
        })
    }
}

// HandleCreateAdmin crea el usuario administrador
func HandleCreateAdmin(store *storage.Storage) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var admin struct {
            Email    string `json:"email"`
            Password string `json:"password"`
        }
        
        if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }
        
        if err := store.CreateAdminUser(admin.Email, admin.Password); err != nil {
            http.Error(w, "Error creating admin user: "+err.Error(), http.StatusInternalServerError)
            return
        }
        
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "message": "Admin user created successfully",
        })
    }
}

// HandleSetupComplete finaliza el proceso de setup
func HandleSetupComplete(store *storage.Storage) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if err := store.CompleteSetup(); err != nil {
            http.Error(w, "Error completing setup: "+err.Error(), http.StatusInternalServerError)
            return
        }
        
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "message": "Setup completed successfully",
        })
    }
}