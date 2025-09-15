package storage

import (

    "os"


    "system_boot/internal/utils"  // Agregar esta importación
)

// Eliminar las funciones duplicadas y usar las de utils

// writeFileWithPerms (ahora usa utils)
func writeFileWithPerms(filename string, data []byte, perm os.FileMode) error {
    return utils.WriteFileWithPerms(filename, data, perm)
}

// encryptData (ahora usa utils)
func encryptData(key []byte, plaintext []byte) ([]byte, []byte, error) {
    return utils.EncryptData(key, plaintext)
}

// decryptData (ahora usa utils)
func decryptData(key []byte, ciphertext []byte, nonce []byte) ([]byte, error) {
    return utils.DecryptData(key, ciphertext, nonce)
}

// encodeBase64 (ahora usa utils)
func encodeBase64(data []byte) string {
    return utils.EncodeBase64(data)
}

// decodeBase64 (ahora usa utils)
func decodeBase64(data string) ([]byte, error) {
    return utils.DecodeBase64(data)
}