package config

import "os"

type Config struct {
    Port        string
    Env         string
    MasterKey   string
    JWTSecret   string
    JWTExpire   string
    FrontendURL string
}

func LoadConfig() Config {
    return Config{
        Port:        getEnv("PORT", "8000"),
        Env:         getEnv("ENV", "development"),
        MasterKey:   getEnv("MASTER_KEY", ""),
        JWTSecret:   getEnv("JWT_SECRET", ""),
        JWTExpire:   getEnv("JWT_EXPIRE", "24h"),
        FrontendURL: getEnv("FRONTEND_URL", "http://localhost:3000"),
    }
}

func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}