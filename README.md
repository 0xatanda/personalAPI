# Personal API - DevOps Stage 1

## Overview
A lightweight Go API deployed with Nginx reverse proxy and systemd.

## Base URL
http://your-server-ip-or-domain

---

## Endpoints

### GET /
```json
{ "message": "API is running" }

### GET /health
```json
{ "message": "healthy" }

### GET /me
```json
{
  "name": "Nafiu Atanda",
  "email": "atanda0x@gmail.com.com",
  "github": "https://github.com/0xatanda"
}

go run main.go