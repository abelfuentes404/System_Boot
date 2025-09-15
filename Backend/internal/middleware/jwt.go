package middleware

import (
    "context"
    "net/http"
    "os"
    "strings"

    "github.com/golang-jwt/jwt"
)

// JWTAuthMiddleware verifica el token JWT
func JWTAuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenString := extractToken(r)
        if tokenString == "" {
            http.Error(w, "Token required", http.StatusUnauthorized)
            return
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("JWT_SECRET")), nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            http.Error(w, "Invalid token claims", http.StatusUnauthorized)
            return
        }

        // Agregar claims al contexto
        ctx := context.WithValue(r.Context(), "claims", claims)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func extractToken(r *http.Request) string {
    bearerToken := r.Header.Get("Authorization")
    if strings.HasPrefix(bearerToken, "Bearer ") {
        return strings.TrimPrefix(bearerToken, "Bearer ")
    }
    return ""
}