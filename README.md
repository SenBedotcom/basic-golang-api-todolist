# Todo List API - Clean Architecture

RESTful API สำหรับจัดการ Todo List ที่พัฒนาด้วย Go และใช้หลักการ Clean Architecture

## 🏗️ Architecture

โปรเจกต์นี้ใช้ Clean Architecture แบ่งเป็น 4 ชั้น:

```
├── cmdapi/                    # Entry point
├── internaldomain/           # Domain Layer (Entities & Interfaces)
├── internalusecase/          # Use Case Layer (Business Logic)
├── internalrepository/       # Infrastructure Layer (Database)
├── internaldeliveryhttp/     # Delivery Layer (HTTP Handlers)
├── config/                    # Configuration
└── pkgdatabase/              # Database utilities
```

### Layers

1. **Domain Layer** (`internaldomain/`)
   - Entities (Todo)
   - Repository Interfaces
   - ไม่มี dependencies ต่อ layer อื่น

2. **Use Case Layer** (`internalusecase/`)
   - Business Logic
   - ใช้งาน Domain Interfaces
   - ไม่รู้จัก Implementation details

3. **Repository Layer** (`internalrepository/`)
   - Database Implementation
   - Implements Domain Interfaces
   - PostgreSQL database operations

4. **Delivery Layer** (`internaldeliveryhttp/`)
   - HTTP Handlers
   - Request/Response handling
   - Route setup

## 🚀 Features

- ✅ CRUD operations สำหรับ Todo items
- ✅ RESTful API design
- ✅ Clean Architecture principles
- ✅ PostgreSQL database
- ✅ Dependency Injection
- ✅ Error handling
- ✅ Docker support

## 📋 Prerequisites

- Go 1.19 or higher
- PostgreSQL 15
- Docker & Docker Compose (optional)

## 🔧 Installation

### 1. Clone Repository

```bash
git clone <repository-url>
cd su-basic-go-api-todo-list
```

### 2. Install Dependencies

```bash
make install
# หรือ
go mod download
go mod tidy
```

### 3. Setup Database

#### ใช้ Docker (แนะนำ)

```bash
make docker-up
```

#### ติดตั้ง PostgreSQL เอง

```bash
# สร้าง database
psql -U postgres
CREATE DATABASE todoapp;
```

### 4. Configuration

สร้างไฟล์ `.env` หรือ set environment variables:

```bash
# Server Configuration
export SERVER_PORT=8080

# Database Configuration
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=todoapp
export DB_SSLMODE=disable
```

## 🎯 Usage

### Start Server

```bash
make run
# หรือ
go run cmdapi/main.go
```

Server จะรันที่ `http://localhost:8080`

### API Endpoints

#### Health Check
```bash
GET /health
```

#### Get All Todos
```bash
GET /api/v1/todos
```

#### Get Todo by ID
```bash
GET /api/v1/todos/{id}
```

#### Create Todo
```bash
POST /api/v1/todos
Content-Type: application/json

{
  "title": "Learn Clean Architecture",
  "description": "Study and implement clean architecture in Go"
}
```

#### Update Todo
```bash
PUT /api/v1/todos/{id}
Content-Type: application/json

{
  "title": "Learn Clean Architecture",
  "description": "Study and implement clean architecture in Go",
  "completed": true
}
```

#### Toggle Todo Completion
```bash
PATCH /api/v1/todos/{id}/toggle
```

#### Delete Todo
```bash
DELETE /api/v1/todos/{id}
```

## 📝 Example Requests

### Create Todo

```bash
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Buy groceries",
    "description": "Milk, eggs, bread"
  }'
```

### Get All Todos

```bash
curl http://localhost:8080/api/v1/todos
```

### Update Todo

```bash
curl -X PUT http://localhost:8080/api/v1/todos/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Buy groceries",
    "description": "Milk, eggs, bread, cheese",
    "completed": true
  }'
```

### Toggle Completion

```bash
curl -X PATCH http://localhost:8080/api/v1/todos/1/toggle
```

### Delete Todo

```bash
curl -X DELETE http://localhost:8080/api/v1/todos/1
```

## 🛠️ Development

### Build

```bash
make build
# Binary จะถูกสร้างที่ bin/todo-api
```

### Run Tests

```bash
make test
```

### Clean Build Files

```bash
make clean
```

## 🐳 Docker Commands

```bash
# Start PostgreSQL
make docker-up

# Stop PostgreSQL
make docker-down

# View logs
make docker-logs

# Complete setup
make setup
```

## 📦 Dependencies

- [gorilla/mux](https://github.com/gorilla/mux) - HTTP router
- [lib/pq](https://github.com/lib/pq) - PostgreSQL driver

## 🏛️ Project Structure

```
.
├── cmdapi/
│   └── main.go                      # Application entry point
├── internaldomain/
│   └── todo.go                      # Todo entity & repository interface
├── internalusecase/
│   └── todo_usecase.go             # Business logic
├── internalrepository/
│   └── todo_repository_postgres.go # PostgreSQL implementation
├── internaldeliveryhttp/
│   ├── todo_handler.go             # HTTP handlers
│   └── router.go                    # Route setup
├── config/
│   └── config.go                    # Configuration management
├── pkgdatabase/
│   └── postgres.go                  # Database connection & schema
├── docker-compose.yml               # Docker Compose configuration
├── Makefile                         # Build commands
├── go.mod                           # Go module file
└── README.md                        # This file
```

## 🎓 Clean Architecture Benefits

1. **Independent of Frameworks** - Business logic ไม่ผูกติดกับ framework
2. **Testable** - Business logic สามารถ test ได้โดยไม่ต้องใช้ UI, Database, Web Server
3. **Independent of UI** - UI สามารถเปลี่ยนได้ง่าย โดยไม่กระทบ business logic
4. **Independent of Database** - สามารถเปลี่ยน database ได้โดยไม่กระทบ business logic
5. **Easy to Maintain** - แต่ละ layer มีหน้าที่ชัดเจน ง่ายต่อการดูแลรักษา

## 📚 Learn More

- [Clean Architecture by Uncle Bob](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Go Documentation](https://go.dev/doc/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)

## 📄 License

MIT License

## 👨‍💻 Author

SenBedotcom
