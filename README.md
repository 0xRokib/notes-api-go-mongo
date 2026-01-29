# Notes API 📝

A small RESTful CRUD API for managing notes built with Go, Gin, and MongoDB.

## 🔧 Features

- Create, read, update, delete notes (CRUD)
- Health check endpoint
- MongoDB-backed storage

## 🚀 Tech stack

- Go (module: `notes-api`, go 1.25.1)
- Gin HTTP framework
- MongoDB (official mongo-driver)

## ⚙️ Requirements

- Go 1.25+ installed
- Running MongoDB instance
- Environment variables: `MONGO_URI`, `MONGO_DB`, `PORT`

Create a `.env` file in the project root (example):

```
MONGO_URI=mongodb://localhost:27017
MONGO_DB=notes_db
PORT=8080
```

## 🧭 Run locally

1. Install dependencies and build/run the app:

- Run directly:

```
go run ./cmd/api
```

- Or build and run:

```
go build -o bin/notes-api ./cmd/api
./bin/notes-api
```

The server listens on `:$PORT` (from env).

## 📦 API Endpoints

- Health
  - GET `/health` — returns health status

- Notes
  - POST `/notes` — create note
    - Request body (JSON):

      ```json
      {
        "title": "Buy milk",
        "content": "Get 2 gallons of milk",
        "pinned": false
      }
      ```

    - Success: HTTP 201, created note JSON (includes `id`, `createdAt`, `updatedAt`)

  - GET `/notes` — list all notes
    - Success: HTTP 200, JSON `{ "notes": [ ... ] }`

  - GET `/notes/:id` — get note by ID
    - Success: HTTP 200, returns the note object

  - PUT `/notes/:id` — update note by ID
    - Request body (JSON): partial or full fields

      ```json
      {
        "title": "Buy almond milk",
        "content": "Get unsweetened almond milk",
        "pinned": true
      }
      ```

    - Success: HTTP 200, returns updated note

  - DELETE `/notes/:id` — delete note by ID
    - Success: HTTP 200, returns `{ "deleted": true, "message": "note deleted successfully" }`

### cURL Examples

- Create

```
curl -X POST http://localhost:8080/notes \
  -H "Content-Type: application/json" \
  -d '{"title":"Test","content":"hello","pinned":false}'
```

- List

```
curl http://localhost:8080/notes
```

- Get

```
curl http://localhost:8080/notes/<id>
```

- Update

```
curl -X PUT http://localhost:8080/notes/<id> -H "Content-Type: application/json" -d '{"title":"Updated"}'
```

- Delete

```
curl -X DELETE http://localhost:8080/notes/<id>
```

## 🗂 Project structure

Top-level layout:

```
cmd/api/main.go          # App entrypoint
internal/config/config.go
internal/db/mongo.go
internal/server/router.go
internal/notes/*         # routes, handler, model, repo
go.mod
```

## 💡 Notes

- Collection used: `notes`
- The handlers validate input and return appropriate HTTP codes
- The app uses `.env` via `github.com/joho/godotenv` so `.env` is expected in dev

## ✅ Contributing

Contributions welcome — open an issue or PR with improvements.

## 📜 License

This project is provided as-is. Use a permissive license (MIT) if you want to reuse it.
