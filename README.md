# Greening India Backend

A Go-based backend service with authentication, JWT, PostgreSQL, and Docker support.

---

## 🚀 Quick Start

### 1. Clone the repo

git clone <your-repo-url>
cd greening-india

### 2. Run the application

docker-compose up --build

---

## 🗄️ Database

* PostgreSQL runs via Docker
* Migrations are applied automatically using Goose

---

## 🔐 Authentication

* Passwords are hashed using bcrypt
* JWT-based authentication (24h expiry)
* Secret is loaded from `.env`

---

## 📡 API Endpoints

### Health Check

GET /health

---

### Register

POST /auth/register

Request:
{
"name": "Aadith",
"email": "[test@example.com](mailto:test@example.com)",
"password": "123456"
}

---

### Login

POST /auth/login

Response:
{
"token": "<JWT_TOKEN>"
}

---

### Get Users (Protected)

GET /users

Header:
Authorization: Bearer <JWT_TOKEN>

---

## 🧪 Example Curl

### Register

curl -X POST http://localhost:8080/auth/register 
-H "Content-Type: application/json" 
-d '{"name":"Test","email":"[test@example.com](mailto:test@example.com)","password":"123456"}'

---

### Login

curl -X POST http://localhost:8080/auth/login 
-H "Content-Type: application/json" 
-d '{"email":"[test@example.com](mailto:test@example.com)","password":"123456"}'

---

### Get Users

curl http://localhost:8080/users 
-H "Authorization: Bearer <token>"

---

## ⚙️ Environment Variables

Create a `.env` file in root:

DB_HOST=db
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=greening_india
DB_PORT=5432

JWT_SECRET=supersecret

---

## 🧰 Tech Stack

* Go (Golang)
* Gin
* GORM
* PostgreSQL
* Goose (migrations)
* Docker

---

## ✅ Notes

* App is fully containerized
* DB retry logic implemented
* Migrations run automatically on startup
* Passwords are securely hashed
