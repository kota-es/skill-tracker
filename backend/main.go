package main

import (
	"backend/containers"
	"backend/controllers"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	dbHost := os.Getenv("POSTGRES_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	sslMode := os.Getenv("POSTGRES_SSL_MODE")
	dbConn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbName, sslMode)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		log.Fatal(err)
		return
	}

	serviceContainer := containers.NewServiceContainer(db)
	userController := controllers.NewUserController(serviceContainer)
	authController := controllers.NewAuthController(serviceContainer)

	http.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	http.HandleFunc("POST /users", userController.PostUser)
	http.HandleFunc("POST /login", authController.Login)
	http.HandleFunc("GET /users/me", authController.Me)

	http.ListenAndServe(":8080", nil)
}
