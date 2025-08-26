package storage

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "io"
    "os"
    "path/filepath"

    "golang.org/x/crypto/bcrypt"
)

// Estructuras para los archivos de configuración
type State struct {
    SetupComplete bool `json:"setupComplete"`
}

type Meta struct {
    Version          string `json:"version"`
    Algorithm        string `json:"algorithm"`
    IV               string `json:"iv"`
    EncryptedDataKey string `json:"encrypted_data_key"`
    CreatedAt        string `json:"created_at"`
}

type PostgresCredentials struct {
    Host     string `json:"host"`
    Port     int    `json:"port"`
    User     string `json:"user"`
    Password string `json:"password"`
    DBName   string `json:"dbname"`
}

type LocalCredentials struct {
    Username     string `json:"username"`
    PasswordHash string `json:"password_hash"`
}

type EncryptedFile struct {
    Nonce      string `json:"nonce"`
    Ciphertext string `json:"ciphertext"`
}

// Constantes para nombres de archivos
const (
    stateFile = "state.json"
    metaFile  = "meta.json"
    dbEncFile = "db.enc"
    authEncFile = "auth.enc"
)

// encryptData cifra datos usando AES-GCM
func encryptData(key []byte, plaintext []byte) ([]byte, []byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, nil, err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, nil, err
    }

    nonce := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, nil, err
    }

    ciphertext := gcm.Seal(nil, nonce, plaintext, nil)
    return ciphertext, nonce, nil
}

// decryptData descifra datos usando AES-GCM
func decryptData(key []byte, ciphertext []byte, nonce []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return nil, err
    }

    return plaintext, nil
}

// writeFileWithPerms escribe un archivo con permisos restringidos
func writeFileWithPerms(filename string, data []byte, perm os.FileMode) error {
    // Asegurar que el directorio existe
    dir := filepath.Dir(filename)
    if err := os.MkdirAll(dir, 0700); err != nil {
        return err
    }
    
    // Escribir archivo
    if err := os.WriteFile(filename, data, perm); err != nil {
        return err
    }
    
    return nil
}

// hashPassword genera un hash bcrypt de una contraseña
func hashPassword(password string) (string, error) {
    hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedBytes), nil
}

// checkPasswordHash verifica una contraseña contra un hash bcrypt
func checkPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}