package storage

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
	"path/filepath"
	"system_boot/internal/utils"
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
// encodeBase64 (ahora usa utils)
func encodeBase64(data []byte) string {
    return utils.EncodeBase64(data)
}

// decodeBase64 (ahora usa utils)
func decodeBase64(data string) ([]byte, error) {
    return utils.DecodeBase64(data)
}