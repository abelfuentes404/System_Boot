package utils

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "io"
)

// EncryptData cifra datos usando AES-GCM
func EncryptData(key []byte, plaintext []byte) ([]byte, []byte, error) {
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

// DecryptData descifra datos usando AES-GCM
func DecryptData(key []byte, ciphertext []byte, nonce []byte) ([]byte, error) {
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

// GenerateRandomKey genera una clave aleatoria de 32 bytes
func GenerateRandomKey() ([]byte, error) {
    key := make([]byte, 32)
    _, err := rand.Read(key)
    if err != nil {
        return nil, err
    }
    return key, nil
}

// EncodeBase64 codifica bytes a base64
func EncodeBase64(data []byte) string {
    return base64.StdEncoding.EncodeToString(data)
}

// DecodeBase64 decodifica base64 a bytes
func DecodeBase64(data string) ([]byte, error) {
    return base64.StdEncoding.DecodeString(data)
}