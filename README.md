# Golang REST API with MySQL & JWT Authentication

A backend REST API project built with Golang using:

- Gorilla Mux
- MySQL
- JWT Authentication
- Middleware
- CRUD Operations
- Password Hashing (bcrypt)

---

# Features

- User Registration
- User Login
- JWT Authentication
- Protected Routes
- MySQL Database Integration
- CRUD APIs
- Middleware Support
- Password Encryption

---

# Tech Stack

- Golang
- Gorilla Mux
- MySQL
- JWT
- bcrypt

---

# Project Structure

```text
RESTAPI-MYSQL/
│
├── main.go
├── go.mod
├── go.sum
│
├── config/
│   └── db.go
│
├── controllers/
│   ├── authController.go
│   └── userController.go
│
├── middleware/
│   └── authMiddleware.go
│
├── models/
│   └── userModel.go
│
├── routes/
│   └── routes.go
│
├── utils/
│   ├── jwt.go
│   └── password.go
│
└── README.md