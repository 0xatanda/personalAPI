# HNG Stage 1 – Personal API (DevOps Track)

## 🌐 Live URL

https://54.163.5.193.nip.io

---

## 📌 Overview

This project is a simple REST API built with Go and deployed on a Linux VPS using Nginx as a reverse proxy. The application runs as a persistent systemd service and serves three endpoints required for HNG Stage 1.

---

## ⚙️ Tech Stack

* Go (Golang)
* Nginx (Reverse Proxy)
* Ubuntu (EC2 VPS)
* systemd (Process Management)
* Let's Encrypt (SSL via Certbot)

---

## 🚀 How to Run Locally

### 1. Clone the repository

```bash
git clone https://github.com/0xatanda/personalAPI.git
cd personalAPI
```

### 2. Run the app

```bash
go run main.go
```

### 3. Access locally

```bash
http://localhost:8080
```

---

## 📡 API Endpoints

### 1. GET /

Returns:

```json
{
  "message": "API is running"
}
```

---

### 2. GET /health

Returns:

```json
{
  "message": "healthy"
}
```

---

### 3. GET /me

Returns:

```json
{
  "name": "Your Full Name",
  "email": "you@example.com",
  "github": "https://github.com/yourusername"
}
```

---

## ⚙️ Deployment Details

* The Go application runs on **port 8080 (internal only)**
* Nginx listens on **ports 80 and 443**
* Public traffic is proxied to the Go app via:

```
http://127.0.0.1:8080
```

---

## 🔐 Security & Performance

* HTTPS enabled via Let's Encrypt
* HTTP → HTTPS redirect (301)
* API responses return:

  * `Content-Type: application/json`
  * HTTP Status `200`
* Response time < 500ms

---

## 🛠️ Process Management

The API runs as a systemd service:

```bash
sudo systemctl status personalAPI
```

Ensures:

* Automatic restart on failure
* Persistence after reboot

---

## 📁 Project Structure

```
.
├── main.go
├── app (compiled binary)
└── README.md
```

---

## ✅ HNG Requirements Covered

* ✔ All endpoints return JSON
* ✔ HTTP 200 responses
* ✔ Reverse proxy with Nginx
* ✔ App not publicly exposed (port 8080 internal)
* ✔ Persistent service using systemd
* ✔ HTTPS enabled
* ✔ Fast response time

---

## 👤 Author

* Name: Nafiu Atanda
* Email: [atanda0x@gmail.com](mailto:atanda0x@gmail.com)
* GitHub: https://github.com/0xatanda

---
