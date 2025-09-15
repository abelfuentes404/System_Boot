package utils

import (
    "os"
    "path/filepath"
)

// WriteFileWithPerms escribe un archivo con permisos restringidos
func WriteFileWithPerms(filename string, data []byte, perm os.FileMode) error {
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

// FileExists verifica si un archivo existe
func FileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

// EnsureDir asegura que el directorio existe
func EnsureDir(dir string) error {
    return os.MkdirAll(dir, 0700)
}