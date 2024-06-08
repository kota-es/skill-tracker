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
	dbConn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbName)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		log.Fatal(err)
		return
	}

	serviceContainer := containers.NewServiceContainer(db)
	userController := controllers.NewUserController(serviceContainer)

	http.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	http.HandleFunc("POST /users", userController.PostUser)

	http.ListenAndServe(":8080", nil)
}
