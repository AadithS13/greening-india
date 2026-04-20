# Greening India Backend

A Go-based backend service with authentication, JWT, PostgreSQL, and Docker support.

---

## 🚀 Run in one command

```bash
docker compose up --build
```

---

## 📦 Setup

### 1. Clone the repo

```bash
git clone https://github.com/<your-username>/greening-india.git
cd greening-india
```

### 2. Create `.env` file

```env
DB_HOST=db
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=greening_india
DB_PORT=5432

JWT_SECRET=supersecret
```

### 3. Start the application

```bash
docker compose up --build
```

---

## 🗄️ Database

* PostgreSQL runs via Docker
* Migrations are applied automatically using Goose on startup

---

## 🔐 Authentication

* Passwords are hashed using bcrypt (cost ≥ 12)
* JWT-based authentication (24-hour expiry)
* JWT secret is loaded from `.env`

---

## 📡 API Endpoints

### Health Check

```
GET /health
```

---

### Register

```
POST /auth/register
```

**Request**

```json
{
  "name": "Aadith",
  "email": "test@example.com",
  "password": "123456"
}
```

---

### Login

```
POST /auth/login
```

**Response**

```json
{
  "token": "<JWT_TOKEN>"
}
```

---

### Get Users (Protected)

```
GET /users
```

**Headers**

```
Authorization: Bearer <JWT_TOKEN>
```

---

## 🧪 Example Curl

### Register

```bash
curl -X POST http://localhost:8080/auth/register \
-H "Content-Type: application/json" \
-d '{"name":"Test","email":"test@example.com","password":"123456"}'
```

---

### Login

```bash
curl -X POST http://localhost:8080/auth/login \
-H "Content-Type: application/json" \
-d '{"email":"test@example.com","password":"123456"}'
```

---

### Get Users

```bash
curl http://localhost:8080/users \
-H "Authorization: Bearer <token>"
```

---

## 🧰 Tech Stack

* Go (Golang)
* Gin
* GORM
* PostgreSQL
* Goose (migrations)
* Docker

---

## 🏗️ Architecture

```
Handler → Service → Repository → Database
```

---

## ✅ Notes

* Fully containerized application
* Automatic DB retry on startup
* Migrations run automatically
* Secure password storage (bcrypt)
* JWT-based authentication
