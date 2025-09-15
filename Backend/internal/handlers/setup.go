package handlers

import (
    "encoding/json"
    "net/http"

    "system_boot/internal/storage"
)

// GetSetupStatus devuelve el estado actual del setup
func GetSetupStatus(w http.ResponseWriter, r *http.Request, store *storage.Storage) {
    status := map[string]interface{}{
        "setupComplete": store.State.SetupComplete,
        "hasDbConfig":   store.HasDBConfig(),
    }
    
    respondWithJSON(w, http.StatusOK, status)
}

// HandleSetupLogin maneja el login durante el setup
func HandleSetupLogin(w http.ResponseWriter, r *http.Request, store *storage.Storage) {
    var credentials struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    
    if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid JSON")
        return
    }
    
    // Verificar credenciales locales
    if !store.ValidateLocalCredentials(credentials.Username, credentials.Password) {
        respondWithError(w, http.StatusUnauthorized, "Invalid credentials")
        return
    }
    
    // Generar token JWT temporal para el setup
    token, err := GenerateSetupToken(credentials.Username)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Error generating token")
        return
    }
    
    respondWithJSON(w, http.StatusOK, map[string]string{
        "token": token,
        "message": "Login successful",
    })
}

// HandleTestConnection prueba la conexión a PostgreSQL
func HandleTestConnection(w http.ResponseWriter, r *http.Request, store *storage.Storage) {
    var config struct {
        Host     string `json:"host"`
        Port     int    `json:"port"`
        User     string `json:"user"`
        Password string `json:"password"`
        DBName   string `json:"dbname"`
    }
    
    if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid JSON")
        return
    }
    
    // Probar conexión
    if err := store.TestPostgresConnection(config.Host, config.Port, config.User, config.Password, config.DBName); err != nil {
        respondWithError(w, http.StatusBadRequest, "Connection failed: "+err.Error())
        return
    }
    
    respondWithJSON(w, http.StatusOK, map[string]string{
        "message": "Connection successful",
    })
}

// HandleSetupConfiguration guarda la configuración de la base de datos
func HandleSetupConfiguration(w http.ResponseWriter, r *http.Request, store *storage.Storage) {
    var config struct {
        Host     string `json:"host"`
        Port     int    `json:"port"`
        User     string `json:"user"`
        Password string `json:"password"`
        DBName   string `json:"dbname"`
    }
    
    if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid JSON")
        return
    }
    
    // Guardar configuración
    if err := store.SaveDBConfiguration(config.Host, config.Port, config.User, config.Password, config.DBName); err != nil {
        respondWithError(w, http.StatusInternalServerError, "Error saving configuration: "+err.Error())
        return
    }
    
    respondWithJSON(w, http.StatusOK, map[string]string{
        "message": "Configuration saved successfully",
    })
}

// HandleCreateTables crea las tablas en la base de datos
func HandleCreateTables(w http.ResponseWriter, r *http.Request, store *storage.Storage) {
    if err := store.CreateTables(); err != nil {
        respondWithError(w, http.StatusInternalServerError, "Error creating tables: "+err.Error())
        return
    }
    
    respondWithJSON(w, http.StatusOK, map[string]string{
        "message": "Tables created successfully",
    })
}

// HandleCreateAdmin crea el usuario administrador
func HandleCreateAdmin(w http.ResponseWriter, r *http.Request, store *storage.Storage) {
    var admin struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    
    if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid JSON")
        return
    }
    
    if err := store.CreateAdminUser(admin.Email, admin.Password); err != nil {
        respondWithError(w, http.StatusInternalServerError, "Error creating admin user: "+err.Error())
        return
    }
    
    respondWithJSON(w, http.StatusOK, map[string]string{
        "message": "Admin user created successfully",
    })
}

// HandleSetupComplete finaliza el proceso de setup
func HandleSetupComplete(w http.ResponseWriter, r *http.Request, store *storage.Storage) {
    if err := store.CompleteSetup(); err != nil {
        respondWithError(w, http.StatusInternalServerError, "Error completing setup: "+err.Error())
        return
    }
    
    respondWithJSON(w, http.StatusOK, map[string]string{
        "message": "Setup completed successfully",
    })
}