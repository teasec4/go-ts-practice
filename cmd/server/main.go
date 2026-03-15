package main

import (
	"database/sql"
	"justforfun/internal/handler"
	"justforfun/internal/repository"
	"justforfun/internal/service"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users(
		id TEXT PRIMARY KEY,
		name TEXT
		)`)
	if err != nil {
		log.Fatal(err)
	}

	var userRepo repository.UserRepository
	userRepo = repository.NewSQLiteUserRepositoryImpl(db)

	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// CORS middleware
	corsMiddleware := func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next(w, r)
		}
	}

	http.HandleFunc("/users", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			userHandler.GetAllUsers(w, r)
		} else if r.Method == http.MethodPost {
			userHandler.CreateUser(w, r)
		}
	}))

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handler.IndexHandler)

	http.HandleFunc("/user", corsMiddleware(userHandler.FindUserByID))

	// Update app.js to use correct path
	http.HandleFunc("/app.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/app.js")
	})

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
