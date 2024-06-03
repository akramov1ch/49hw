package main

import (
	"fmt"
	"log"
	"net/http"
	"49hw/config"
	"49hw/db"
	"49hw/handlers"
	"49hw/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.LoadConfig()
	dbConn, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	db.RunMigrations(dbConn)

	userService := &services.UserService{DB: dbConn}
	userHandler := &handlers.UserHandler{UserService: userService}

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/register", userHandler.Register)
	router.Post("/login", userHandler.Login)

	fmt.Println("Starting the server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
