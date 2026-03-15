# Blog REST API

A REST API for a blog application built with Golang and MySQL.

## Tech Stack
- **Golang** — Backend language
- **Gin** — Web framework
- **MySQL** — Database
- **GORM** — ORM
- **JWT** — Authentication
- **bcrypt** — Password hashing

## Project Structure
```
blog-api/
├── main.go
├── .env
├── config/
│   └── database.go
├── models/
│   └── models.go
├── controllers/
│   ├── auth.go
│   └── post.go
└── routes/
    ├── routes.go
    └── middleware.go
```

## Setup

### 1. Database banao
```sql
CREATE DATABASE blogdb;
```

### 2. .env file banao
```
DB_USER=root
DB_PASSWORD=yourpassword
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=blogdb
JWT_SECRET=your-secret-key
```

### 3. Dependencies install karo
```bash
go mod tidy
```

### 4. Run karo
```bash
go run main.go
```

## API Endpoints

### Public
| Method | URL | Description |
|--------|-----|-------------|
| POST | /api/register | Register new user |
| POST | /api/login | Login and get token |
| GET | /api/posts | Get all posts |
| GET | /api/posts/:id | Get single post |

### Protected (Token required)
| Method | URL | Description |
|--------|-----|-------------|
| POST | /api/posts | Create post |
| PUT | /api/posts/:id | Update post |
| DELETE | /api/posts/:id | Delete post |

## Author
**Rupendra Kumar** — Golang Backend Developer
GitHub: [ReyanshGit](https://github.com/ReyanshGit)
