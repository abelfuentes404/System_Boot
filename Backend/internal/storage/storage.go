package storage

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type Storage struct {
    State           *State
    postgresCreds   *PostgresCredentials
    localCreds      *LocalCredentials
    db              *sql.DB
}

func NewStorage() *Storage {
    return &Storage{
        State: &State{SetupComplete: false},
    }
}

// LoadState carga el estado y configuraciones desde archivos
func (s *Storage) LoadState() error {
    // Cargar state.json
    if stateData, err := os.ReadFile(stateFile); err == nil {
        json.Unmarshal(stateData, &s.State)
    }

    // Si el setup está completo, cargar configuraciones
    if s.State.SetupComplete {
        if err := s.loadConfigurations(); err != nil {
            // Si hay error al cargar configuraciones, resetear
            log.Printf("Error loading configurations: %v", err)
            s.Reset()
            return err
        }
    }

    return nil
}

// loadConfigurations carga las configuraciones cifradas
func (s *Storage) loadConfigurations() error {
    // Cargar meta.json
    metaData, err := os.ReadFile(metaFile)
    if err != nil {
        return err
    }
    
    var meta Meta
    if err := json.Unmarshal(metaData, &meta); err != nil {
        return err
    }
    
    // Obtener master key de environment
    masterKey := os.Getenv("MASTER_KEY")
    if masterKey == "" {
        return fmt.Errorf("MASTER_KEY not set")
    }
    
    // Descifrar data key
    encryptedDataKey, err := base64.StdEncoding.DecodeString(meta.EncryptedDataKey)
    if err != nil {
        return err
    }
    
    nonce, err := base64.StdEncoding.DecodeString(meta.IV)
    if err != nil {
        return err
    }
    
    dataKey, err := decryptData([]byte(masterKey), encryptedDataKey, nonce)
    if err != nil {
        s.Reset()
        return fmt.Errorf("failed to decrypt data key: %v", err)
    }
    
    // Descifrar db.enc
    dbEncData, err := os.ReadFile(dbEncFile)
    if err != nil {
        return err
    }
    
    var dbEnc EncryptedFile
    if err := json.Unmarshal(dbEncData, &dbEnc); err != nil {
        return err
    }
    
    encryptedPostgres, err := base64.StdEncoding.DecodeString(dbEnc.Ciphertext)
    if err != nil {
        return err
    }
    
    noncePostgres, err := base64.StdEncoding.DecodeString(dbEnc.Nonce)
    if err != nil {
        return err
    }
    
    postgresCredsBytes, err := decryptData(dataKey, encryptedPostgres, noncePostgres)
    if err != nil {
        s.Reset()
        return fmt.Errorf("failed to decrypt postgres credentials: %v", err)
    }
    
    if err := json.Unmarshal(postgresCredsBytes, &s.postgresCreds); err != nil {
        return err
    }
    
    // Descifrar auth.enc
    authEncData, err := os.ReadFile(authEncFile)
    if err != nil {
        return err
    }
    
    var authEnc EncryptedFile
    if err := json.Unmarshal(authEncData, &authEnc); err != nil {
        return err
    }
    
    encryptedAuth, err := base64.StdEncoding.DecodeString(authEnc.Ciphertext)
    if err != nil {
        return err
    }
    
    nonceAuth, err := base64.StdEncoding.DecodeString(authEnc.Nonce)
    if err != nil {
        return err
    }
    
    authCredsBytes, err := decryptData(dataKey, encryptedAuth, nonceAuth)
    if err != nil {
        s.Reset()
        return fmt.Errorf("failed to decrypt auth credentials: %v", err)
    }
    
    if err := json.Unmarshal(authCredsBytes, &s.localCreds); err != nil {
        return err
    }
    
    return nil
}

// TestPostgresConnection prueba la conexión a PostgreSQL
func (s *Storage) TestPostgresConnection(host string, port int, user string, password string, dbname string) error {
    connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
        host, port, user, password, dbname)
    
    db, err := sql.Open("postgres", connectionString)
    if err != nil {
        return err
    }
    defer db.Close()
    
    return db.Ping()
}

// SaveDBConfiguration guarda la configuración de la base de datos
func (s *Storage) SaveDBConfiguration(host string, port int, user string, password string, dbname string) error {
    s.postgresCreds = &PostgresCredentials{
        Host:     host,
        Port:     port,
        User:     user,
        Password: password,
        DBName:   dbname,
    }
    
    return nil
}

// ValidateLocalCredentials valida las credenciales locales
func (s *Storage) ValidateLocalCredentials(username, password string) bool {
    if s.localCreds == nil {
        // Durante el setup inicial, usar credenciales por defecto
        defaultUser := "admin"
        defaultPass := "admin123"
        
        if username == defaultUser && password == defaultPass {
            return true
        }
        return false
    }
    
    return username == s.localCreds.Username && checkPasswordHash(password, s.localCreds.PasswordHash)
}

// CreateTables crea las tablas en la base de datos
func (s *Storage) CreateTables() error {
    if s.postgresCreds == nil {
        return fmt.Errorf("database configuration not set")
    }
    
    connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
        s.postgresCreds.Host, s.postgresCreds.Port, s.postgresCreds.User, 
        s.postgresCreds.Password, s.postgresCreds.DBName)
    
    db, err := sql.Open("postgres", connectionString)
    if err != nil {
        return err
    }
    defer db.Close()
    
    // Crear tablas
    tablesSQL := `
        CREATE TABLE IF NOT EXISTS smtp_config (
            id SERIAL PRIMARY KEY,
            host VARCHAR(100) NOT NULL,
            port INTEGER NOT NULL,
            username VARCHAR(100) NOT NULL,
            password VARCHAR(100) NOT NULL,
            from_email VARCHAR(100) NOT NULL,
            is_active BOOLEAN DEFAULT TRUE,
            created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
            updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
        );
        
        CREATE INDEX IF NOT EXISTS idx_smtp_config_active ON smtp_config(is_active);
        
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            email VARCHAR(100) UNIQUE NOT NULL,
            password VARCHAR(100) NOT NULL,
            created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
            updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
        );
        
        CREATE TABLE IF NOT EXISTS reset_codes (
            id SERIAL PRIMARY KEY,
            user_id INTEGER NOT NULL,
            code VARCHAR(10) NOT NULL,
            expiration_time TIMESTAMP WITH TIME ZONE NOT NULL,
            created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
            FOREIGN KEY (user_id) REFERENCES users(id)
        );
        
        CREATE INDEX IF NOT EXISTS idx_reset_codes_expiration ON reset_codes(expiration_time);
    `
    
    _, err = db.Exec(tablesSQL)
    return err
}

// CreateAdminUser crea el usuario administrador
func (s *Storage) CreateAdminUser(email, password string) error {
    if s.postgresCreds == nil {
        return fmt.Errorf("database configuration not set")
    }
    
    connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
        s.postgresCreds.Host, s.postgresCreds.Port, s.postgresCreds.User, 
        s.postgresCreds.Password, s.postgresCreds.DBName)
    
    db, err := sql.Open("postgres", connectionString)
    if err != nil {
        return err
    }
    defer db.Close()
    
    // Hash de la contraseña
    hashedPassword, err := hashPassword(password)
    if err != nil {
        return err
    }
    
    // Insertar usuario administrador
    _, err = db.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", email, hashedPassword)
    return err
}

// CompleteSetup finaliza el proceso de setup
func (s *Storage) CompleteSetup() error {
    // Generar data key
    dataKey := make([]byte, 32)
    if _, err := rand.Read(dataKey); err != nil {
        return err
    }
    
    // Obtener master key
    masterKey := os.Getenv("MASTER_KEY")
    if masterKey == "" {
        return fmt.Errorf("MASTER_KEY not set")
    }
    
    // Cifrar credenciales de PostgreSQL
    postgresCredsBytes, err := json.Marshal(s.postgresCreds)
    if err != nil {
        return err
    }
    
    encryptedPostgres, noncePostgres, err := encryptData(dataKey, postgresCredsBytes)
    if err != nil {
        return err
    }
    
    // Guardar db.enc
    dbEnc := EncryptedFile{
        Nonce:      base64.StdEncoding.EncodeToString(noncePostgres),
        Ciphertext: base64.StdEncoding.EncodeToString(encryptedPostgres),
    }
    
    dbEncBytes, err := json.Marshal(dbEnc)
    if err != nil {
        return err
    }
    
    if err := writeFileWithPerms(dbEncFile, dbEncBytes, 0600); err != nil {
        return err
    }
    
    // Cifrar data key con master key
    encryptedDataKey, nonceDataKey, err := encryptData([]byte(masterKey), dataKey)
    if err != nil {
        return err
    }
    
    // Guardar meta.json
    meta := Meta{
        Version:          "1.0",
        Algorithm:        "aes-256-gcm",
        IV:               base64.StdEncoding.EncodeToString(nonceDataKey),
        EncryptedDataKey: base64.StdEncoding.EncodeToString(encryptedDataKey),
        CreatedAt:        time.Now().Format(time.RFC3339),
    }
    
    metaBytes, err := json.Marshal(meta)
    if err != nil {
        return err
    }
    
    if err := writeFileWithPerms(metaFile, metaBytes, 0600); err != nil {
        return err
    }
    
    // Crear credenciales locales (si no existen)
    if s.localCreds == nil {
        hashedPassword, err := hashPassword("admin123") // Contraseña por defecto
        if err != nil {
            return err
        }
        
        s.localCreds = &LocalCredentials{
            Username:     "admin",
            PasswordHash: hashedPassword,
        }
    }
    
    // Cifrar y guardar credenciales locales
    localCredsBytes, err := json.Marshal(s.localCreds)
    if err != nil {
        return err
    }
    
    encryptedAuth, nonceAuth, err := encryptData(dataKey, localCredsBytes)
    if err != nil {
        return err
    }
    
    authEnc := EncryptedFile{
        Nonce:      base64.StdEncoding.EncodeToString(nonceAuth),
        Ciphertext: base64.StdEncoding.EncodeToString(encryptedAuth),
    }
    
    authEncBytes, err := json.Marshal(authEnc)
    if err != nil {
        return err
    }
    
    if err := writeFileWithPerms(authEncFile, authEncBytes, 0600); err != nil {
        return err
    }
    
    // Marcar setup como completo
    s.State.SetupComplete = true
    stateBytes, err := json.Marshal(s.State)
    if err != nil {
        return err
    }
    
    if err := writeFileWithPerms(stateFile, stateBytes, 0600); err != nil {
        return err
    }
    
    return nil
}

// Reset elimina todos los archivos de configuración
func (s *Storage) Reset() {
    os.Remove(stateFile)
    os.Remove(metaFile)
    os.Remove(dbEncFile)
    os.Remove(authEncFile)
    s.State.SetupComplete = false
    s.postgresCreds = nil
    s.localCreds = nil
}

// HasDBConfig verifica si hay configuración de base de datos
func (s *Storage) HasDBConfig() bool {
    return s.postgresCreds != nil
}