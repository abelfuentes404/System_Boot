package handlers

import (
    "encoding/json"
    "net/http"
    "os"
    "time"

    "github.com/golang-jwt/jwt"

    "system_boot/internal/storage"
)

// GenerateSetupToken genera un token JWT para el setup
func GenerateSetupToken(username string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": username,
        "exp":      time.Now().Add(time.Hour * 1).Unix(),
        "purpose":  "setup",
    })

    return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// GenerateAuthToken genera un token JWT para autenticación normal
func GenerateAuthToken(userID int, email string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "userID": userID,
        "email":  email,
        "exp":    time.Now().Add(time.Hour * 24).Unix(),
    })

    return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// HandleAuthLogin maneja el login de usuarios
func HandleAuthLogin(w http.ResponseWriter, r *http.Request, store *storage.Storage) {
    var credentials struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid JSON")
        return
    }

    // Verificar credenciales en la base de datos
    user, err := store.ValidateUserCredentials(credentials.Email, credentials.Password)
    if err != nil || user == nil {
        respondWithError(w, http.StatusUnauthorized, "Invalid credentials")
        return
    }

    // Generar token JWT
    token, err := GenerateAuthToken(user.ID, user.Email)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Error generating token")
        return
    }

    respondWithJSON(w, http.StatusOK, map[string]interface{}{
        "token": token,
        "user": map[string]interface{}{
            "id":    user.ID,
            "email": user.Email,
        },
    })
}

// HandleVerifyToken verifica si un token es válido
func HandleVerifyToken(w http.ResponseWriter, r *http.Request, store *storage.Storage) {
    tokenString := r.Header.Get("Authorization")
    if tokenString == "" {
        respondWithError(w, http.StatusUnauthorized, "Token required")
        return
    }

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("JWT_SECRET")), nil
    })

    if err != nil || !token.Valid {
        respondWithError(w, http.StatusUnauthorized, "Invalid token")
        return
    }

    respondWithJSON(w, http.StatusOK, map[string]interface{}{
        "valid": true,
        "user": token.Claims.(jwt.MapClaims),
    })
}

// HandleAuthLogout cierra la sesión
func HandleAuthLogout(w http.ResponseWriter, r *http.Request) {
    respondWithJSON(w, http.StatusOK, map[string]string{"message": "Logged out successfully"})
}

// HandleDashboardData devuelve datos del dashboard
func HandleDashboardData(w http.ResponseWriter, r *http.Request, store *storage.Storage) {
    respondWithJSON(w, http.StatusOK, map[string]interface{}{
        "message": "Welcome to the dashboard",
        "stats": map[string]int{
            "users": 15,
            "items": 42,
            "sales": 1234,
        },
    })
}