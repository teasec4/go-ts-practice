# JustSome - Go + Hyperapp Fullstack Application

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://golang.org/)
[![Hyperapp](https://img.shields.io/badge/Hyperapp-2.0.22-FF6600)](https://hyperapp.dev/)
[![SQLite](https://img.shields.io/badge/SQLite-3-003B57?logo=sqlite)](https://sqlite.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

A lightweight fullstack application with Go backend and Hyperapp frontend.

## Project Structure

```
justsome/
├── cmd/server/           # Go server entry point
│   └── main.go          # HTTP server with CORS and routing
├── internal/            # Go internal packages
│   ├── handler/         # HTTP handlers
│   ├── model/           # Data models
│   ├── repository/      # Data access layer
│   └── service/         # Business logic
├── static/              # Frontend assets
│   ├── index.html       # Main HTML page
│   ├── app.js           # Hyperapp frontend application
│   ├── api.js           # API client for frontend
│   └── app_example.js   # Example Hyperapp code
└── users.db             # SQLite database
```

## Technologies

### Backend
- **Go** - HTTP server and business logic
- **SQLite** - Lightweight database
- **Go SQLite3 Driver** - Database connectivity
- **Standard net/http** - HTTP server with CORS middleware

### Frontend
- **Hyperapp** - Minimal frontend framework
- **ES Modules** - Modern JavaScript modules
- **Fetch API** - HTTP requests to backend

## Features
- RESTful API for user management
- CORS-enabled for frontend-backend communication
- SQLite database with automatic table creation
- Clean architecture with separation of concerns
- Minimal JavaScript frontend with Hyperapp

## Quick Start

1. **Clone and build:**
   ```bash
   go build -o server cmd/server/main.go
   ```

2. **Run the server:**
   ```bash
   ./server
   ```

3. **Open in browser:**
   ```
   http://localhost:8080/
   ```

## API Endpoints
- `GET /users` - Get all users
- `POST /users` - Create new user
- `GET /user?id={id}` - Find user by ID

## Development
The frontend uses ES modules loaded directly from CDN. Backend follows Go best practices with clean architecture patterns.

## License
MIT