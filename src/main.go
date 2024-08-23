package main

import (
	"log"
	"net/http"
	"time"

	"github.com/themichaelb/Go-WebDev-Learning/internal/config"
	"github.com/themichaelb/Go-WebDev-Learning/internal/routes"
)

func main() {
	const ReadHeaderTimeout = 5 * time.Second
	const ReadTimeout = 5 * time.Second
	const WriteTimeout = 10 * time.Second
	const IdleTimeout = 15 * time.Second

	config.LoadConfig()

	router := routes.SetupRoutes()

	log.Println("Server starting on port 8080...")
	server := &http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadHeaderTimeout: ReadHeaderTimeout,
		ReadTimeout:       ReadTimeout,
		WriteTimeout:      WriteTimeout,
		IdleTimeout:       IdleTimeout,
	}

	log.Println("Server starting on port 8080...")
	log.Fatal(server.ListenAndServe())
}
