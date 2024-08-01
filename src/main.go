package main

import (
	"log"
	"net/http"
	"time"

	"github.com/themichaelb/Go-WebDev-Learning/internal/config"
	"github.com/themichaelb/Go-WebDev-Learning/internal/routes"
)

func main() {
	config.LoadConfig()


	log.Println("Server starting on port 8080...")
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    15 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}
	
	router := routes.SetupRoutes()
	
	log.Println("Server starting on port 8080...")
	log.Fatal(server.ListenAndServe())
}
