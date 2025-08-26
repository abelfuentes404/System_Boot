# 🚀 System Boot

[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go\&logoColor=white)](https://go.dev/)
[![React](https://img.shields.io/badge/React-18-61DAFB?logo=react\&logoColor=black)](https://react.dev/)
[![License](https://img.shields.io/badge/License-Apache%202.0-green.svg)](LICENSE)
[![Build](https://img.shields.io/badge/Build-Passing-brightgreen?logo=githubactions\&logoColor=white)](../../actions)
[![Status](https://img.shields.io/badge/Status-Active-success)]()
[![Made with Love](https://img.shields.io/badge/Made%20with-💙-blue)]()

---

**System Boot** es un sistema fullstack modular para inicializar aplicaciones de manera **segura y automática**.

* 🐹 **Backend en Go**
* ⚛️ **Frontend en React**
* ⚙️ Transición automática entre **modo configuración** y **modo producción**.

---

## 📌 Características principales

* 🔐 Manejo seguro de credenciales usando **cifrado autenticado**.
* 🔄 Transición automática entre setup y producción según `setupComplete`.
* 🛡️ Integración con PostgreSQL (productos) y MongoDB (usuarios/ventas).
* 🌐 Arquitectura fullstack: **Go + React**.
* 🛠️ Reinicio seguro: reset elimina secretos y regresa a setup.

---

## 🔑 Arquitectura de Configuración y Producción

### Estado del sistema

* `setupComplete = false` → **modo configuración** (setup inicial).
* `setupComplete = true` → **modo producción**.
* **Reset:** vuelve a `false` y elimina los secretos.

### Separación de archivos

* `state.json` → estado y metadatos (sin secretos).
* `db.enc` → credenciales de PostgreSQL cifradas.
* `meta.json` → info de cifrado (algoritmo, salt, versión, fechas).

### Protección de secretos

* Nunca guardar JSON plano con credenciales.
* Cifrado autenticado: **AES-256-GCM** o **XChaCha20-Poly1305**.
* **Envelope encryption:**

  1. Data key aleatoria cifra credenciales.
  2. Master key (env, keyring o secret manager) cifra la data key.

### Flujo

1. **Setup:** admin ingresa datos → app valida → cifra y guarda `db.enc` → `setupComplete = true`.
2. **Producción:** al arrancar, descifra secretos en memoria → arma la cadena de conexión → conecta servicios.
3. **Reset:** elimina `db.enc` y deja `setupComplete = false`.

### 🔄 Flujo Visual del Sistema

```
        ┌──────────────┐
        │  Inicio App  │
        └──────┬───────┘
               │
               ▼
      ┌─────────────────┐
      │ setupComplete?  │
      └──────┬──────────┘
       false │ true
             │
             ▼
   ┌───────────────────┐
   │   Modo Producción │
   │ Descifra secretos │
   │ Conecta servicios │
   └───────────────────┘
       ▲
       │
       │
┌───────────────┐
│ Modo Configuración │
│ Ingresa datos      │
│ Valida y cifra     │
│ Guarda db.enc      │
└───────────────┘
       │
       ▼
  setupComplete = true
       │
       ▼
    Repetir flujo
       │
       ▼
      Reset?
       │
       ▼
┌───────────────┐
│ Elimina db.enc│
│ setupComplete=│
│ false         │
└───────────────┘
```

### Rotación y seguridad

* Rotar data key y credenciales periódicamente.
* No loguear secretos, solo eventos (setup, reset, rotación).
* Permisos estrictos en archivos (600) y directorios (700).
* Auditoría de integridad: si falla el descifrado → no arranca, vuelve a setup.

---

## ⚙️ Backend (Go)

📂 Carpeta: `/backend`

* Ejecuta el backend con `go run main.go`.
* La bandera `setupComplete` determina el modo automáticamente.

Endpoints principales:

* `POST /config` → guarda configuración inicial.
* `GET /status` → retorna estado del sistema.
* `GET /db/connection` → devuelve cadena de conexión dinámica.

---

## ⚛️ Frontend (React)

📂 Carpeta: `/frontend`

* Ejecuta con `npm start`.
* Pantalla inicial depende de `setupComplete`:

  * **false:** formulario de setup.
  * **true:** dashboard productivo conectado al backend.

---

## 📦 Instalación rápida

1. Clona el repositorio:

```bash
git clone https://github.com/abelheddy/System_Boot.git
cd system-boot
```

2. Frontend:

```bash
cd frontend
npm install
npm start
```

3. Backend:

```bash
cd ../backend
go mod tidy
go run main.go
```

---

## 🛰️ Roadmap

* Cifrado AES/XChaCha20 completo.
* Dashboard React avanzado para gestión de secretos.
* Contenedores Docker (Go + React).
* CI/CD con GitHub Actions.
* Monitor de rotación de data key.

---

## 📜 Licencia

Este proyecto está licenciado bajo los términos de la Apache License 2.0.
Copyright © 2025 Abel Fuentes Guzman.

---

## 👨‍💻 Autor

**Abel Fuentes Guzman**
🐙 [GitHub](https://github.com/abelheddy)
📧 Contacto: [abelfuentes404@gmail.com](mailto:abelfuentes404@gmail.com)
