package main

import (
	"database/sql"
	"justforfun/internal/handler"
	"justforfun/internal/repository"
	"justforfun/internal/service"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()
	
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users(
		id TEXT PRIMARY KEY,
		name TEXT
		)`)
	if err != nil{
		log.Fatal(err)
	}
	
	var userRepo repository.UserRepository
	userRepo = repository.NewSQLiteUserRepositoryImpl(db)
	
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet{
			userHandler.GetAllUsers(w, r)
		} else if r.Method == http.MethodPost{
			userHandler.CreateUser(w, r)
		}
	})
	
	http.HandleFunc("/user", userHandler.FindUserByID)
	
	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
