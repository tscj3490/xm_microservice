package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/tscj3490/xm_microservice/handlers"
)

func main() {

	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the PostgreSQL database
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	// Create new Router
	router := mux.NewRouter()
	router.Handle("/companies", handlers.CreateCompanyHandler()).Methods("POST")
	router.Handle("/companies/{id}", handlers.GetCompanyHandler()).Methods("GET")
	router.Handle("/companies/{id}", handlers.DeleteCompanyHandler()).Methods("DELETE")
	router.Handle("/companies/{id}", handlers.UpdateCompanyHandler()).Methods("PUT")

	// Create new server and assign the router
	server := http.Server{
		Addr:    ":9090",
		Handler: handlers.AuthHandler(router),
	}

	fmt.Println("Staring Company Microservice on Port 9090")
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to start HTTP Server")
	}
}
