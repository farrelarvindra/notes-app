# 📝 Notes App

A modern, containerized full-stack note-taking application built with **Go**, **PostgreSQL**, **React 19**, **Tailwind CSS v4**, and **Docker Compose**.

---

## 🚀 Tech Stack

### Backend
- **Language**: Go (Golang)
- **HTTP Routing**: Go Standard Library `net/http` (`http.ServeMux` with pattern matching)
- **ORM**: [GORM](https://gorm.io/) with automatic schema migrations
- **Database Driver**: `gorm.io/driver/postgres`
- **Environment**: `godotenv`

### Frontend
- **Framework**: React 19 + TypeScript
- **Bundler & Tooling**: Vite
- **Styling**: Tailwind CSS v4 (`@tailwindcss/vite`)
- **Web Server / Reverse Proxy**: Nginx (Alpine)

### Database & DevOps
- **Database**: PostgreSQL 16 (Alpine)
- **Containers**: Docker with multi-stage builds
- **Orchestration**: Docker Compose

---

## 📁 Project Structure

```text
notes-app/
├── backend/
│   ├── handler/         # API endpoints, request validation, CORS middleware
│   │   └── handler.go
│   ├── model/           # GORM data models (Note struct)
│   │   └── note.go
│   ├── store/           # Database connection & auto-migration
│   │   └── db.go
│   ├── Dockerfile       # Multi-stage build (golang:alpine -> alpine)
│   ├── go.mod           # Go module definitions
│   └── main.go          # Server configuration & graceful shutdown
├── frontend/
│   ├── src/
│   │   ├── components/  # Reusable UI components (Navbar, CreateNoteModal)
│   │   ├── App.tsx      # Main application logic & state management
│   │   ├── main.tsx     # React entrypoint
│   │   └── index.css    # Tailwind CSS imports
│   ├── nginx.conf       # SPA routing & /api reverse proxy configuration
│   ├── Dockerfile       # Multi-stage build (node:alpine -> nginx:alpine)
│   ├── package.json
│   └── vite.config.ts
├── docker-compose.yaml  # Container orchestration (db, backend, frontend)
├── .env.example         # Environment template
├── .gitignore
└── README.md
```

---

## ⚡ Quick Start (Docker Compose)

The fastest and recommended way to run the entire application stack:

### 1. Clone & Setup Environment
```bash
git clone https://github.com/farrelarvindra/notes-app.git
cd notes-app

# Copy environment variables
cp .env.example .env
```

### 2. Run with Docker Compose
```bash
docker compose up -d --build
```

### 3. Access the Application
- **Frontend Web UI**: [http://localhost:3000](http://localhost:3000)
- **Backend API Direct**: [http://localhost:8080/api/](http://localhost:8080/api/)
- **PostgreSQL Database**: `localhost:5432`

---

## 🛠️ Local Development

If you prefer running services directly on your host machine for active development:

### 1. Start Database
```bash
docker compose up -d db
```

### 2. Run Backend
```bash
cd backend
go run main.go
```
*The backend server will run on `http://localhost:8080`.*

### 3. Run Frontend
```bash
cd frontend
npm install
npm run dev
```
*The Vite development server will run on `http://localhost:5173` with proxying to backend.*

---

## 🔌 REST API Endpoints

All API endpoints are prefixed with `/api`.

| Method | Endpoint | Description | Sample Request Body | Status Codes |
|---|---|---|---|---|
| `GET` | `/api/health` | Health check endpoint | _None_ | `200 OK` |
| `GET` | `/api/notes` | Get all notes (newest first) | _None_ | `200 OK` |
| `POST` | `/api/notes` | Create a new note | `{"title": "Note Title", "content": "Note details"}` | `201 Created`, `400 Bad Request` |
| `PUT` | `/api/notes/{id}` | Update an existing note | `{"title": "Updated Title", "content": "Updated details"}` | `200 OK`, `404 Not Found`, `400 Bad Request` |
| `DELETE` | `/api/notes/{id}` | Delete a note by ID | _None_ | `200 OK`, `404 Not Found` |

### Example cURL Commands

#### Create a Note
```bash
curl -X POST http://localhost:3000/api/notes \
  -H "Content-Type: application/json" \
  -d '{"title": "Belajar Docker", "content": "Membuat fullstack container dengan Go dan React"}'
```

#### Get All Notes
```bash
curl http://localhost:3000/api/notes
```

#### Update a Note
```bash
curl -X PUT http://localhost:3000/api/notes/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Belajar Docker (Selesai)", "content": "Berhasil menghubungkan seluruh container"}'
```

#### Delete a Note
```bash
curl -X DELETE http://localhost:3000/api/notes/1
```

---

## ⚙️ Environment Variables

The application is configured using variables in `.env`:

| Variable | Description | Default |
|---|---|---|
| `POSTGRES_USER` | PostgreSQL superuser username | `user` |
| `POSTGRES_PASSWORD` | PostgreSQL password | `password` |
| `POSTGRES_DB` | Initial database name | `notes` |
| `DATABASE_URL` | PostgreSQL connection URL | `postgres://user:password@localhost:5432/notes?sslmode=disable` |
| `PORT` | Backend HTTP server port | `8080` |
| `FRONTEND_PORT` | Frontend host port in Docker Compose | `3000` |

---

## 📄 License
This project is open-source and available under the [MIT License](LICENSE).

