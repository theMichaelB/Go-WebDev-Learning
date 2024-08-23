package main

import (
	"log"
	"net/http"
	"time"

	"github.com/themichaelb/Go-WebDev-Learning/internal/config"
)

func main() {
	// const ReadHeaderTimeout = 5 * time.Second
	const ReadTimeout = 5 * time.Second
	const WriteTimeout = 10 * time.Second
	// const IdleTimeout = 15 * time.Second

	config.LoadConfig()

	// router := routes.SetupRoutes()

	log.Println("Server starting on port 8080...")
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  ReadTimeout,
		WriteTimeout: WriteTimeout,
	}
	log.Println("Server starting on port 8080...")
	log.Fatal(server.ListenAndServe())
}
