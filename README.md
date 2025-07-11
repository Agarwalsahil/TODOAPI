
# 📝 TODOAPI – RESTful ToDo API in Go

![Go CI](https://github.com/Agarwalsahil/TODOAPI/actions/workflows/go.yml/badge.svg)

A secure and efficient TODO list REST API built with **Go**, using **Gin**, **GORM**, **SQLite**, and **JWT authentication**. This is my project following [Todo List API](https://roadmap.sh/projects/todo-list-api).

---

## 🚀 Features

- ✅ User registration and login with hashed passwords.
- ✅ JWT-based authentication middleware.
- ✅ CRUD operations for Todo items.
- ✅ Pagination and filtering on todo list.
- ✅ Per-user scoped data access.
- ✅ SQLite database (simple and lightweight).
- ✅ Clean project structure with modular code.

---

## 📁 Folder Structure

```
TODOAPI/
├── controllers/        # Request handlers
├── db/                 # DB initialization
├── middleware/         # JWT auth middleware
├── models/             # DB models
├── utils/              # Shared helper functions
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
├── main.go
└── README.md
```

---

## 🔧 Setup Instructions

### 1. 📦 Clone the Repository

```bash
git clone https://github.com/Agarwalsahil/TODOAPI.git
cd TODOAPI
```

### 2. ✅ Install Dependencies

```bash
go mod tidy
```

### 3. ▶️ Run the API

```bash
go run main.go
```

Server will run at `http://localhost:8080`

---

## 🛠️ API Endpoints

### 🔐 Authentication

#### 🧾 Register
```http
POST /register
```

**Body:**
```json
{
  "name": "your name",
  "email": "user@example.com",
  "password": "securepassword"
}
```

#### 🔑 Login
```http
POST /login
```

**Body:**
```json
{
  "name": "your name",
  "email": "user@example.com",
  "password": "securepassword"
}
```

Returns:
```json
{
  "token": "<jwt_token>"
}
```

---

### 📋 Todo Operations (🔐 Require JWT in `Authorization` header)

#### ➕ Create Todo
```http
POST /todos
```

**Body:**
```json
{
  "title": "Learn Go",
  "description": "Finish GORM chapter"
}
```

#### 📄 List Todos with Pagination & Filtering
```http
GET /todos?page=1&limit=5
```

#### ✏️ Update Todo
```http
PUT /todos/:id
```

**Body:**
```json
{
  "title": "Updated title",
  "description": "Updated description"
}
```

#### 🗑️ Delete Todo
```http
DELETE /todos/:id
```

---

## 📦 Example Curl Usage

### Register:
```bash
curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{"name": "your name", "email": "user@example.com", "password": "pass"}'
```

### Login:
```bash
curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"name": "your name", "email": "user@example.com", "password": "pass"}'
```

### Create Todo:
```bash
curl -X POST http://localhost:8080/todos -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d '{"title": "New Task", "description": "Do it now"}'
```

---

## 🧾 Technologies Used

- **Go** – Programming language
- **Gin** – HTTP web framework
- **GORM** – ORM for database
- **SQLite** – Lightweight database
- **JWT** – Token-based authentication
- **BCrypt** – Password hashing

---

## 🛡️ Security Notes

- Passwords are hashed using bcrypt before storing.
- JWT secret should be kept in a `.env` file (not committed).
- Authorization middleware protects all todo routes.

---

## 📃 License

This project is licensed under the [MIT License](./LICENSE).

---

## 🙌 Contributions

Open to suggestions, bug reports, and pull requests!

---

Made with ❤️ by [Sahil Agarwal](https://github.com/Agarwalsahil)
