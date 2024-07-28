package main

import (
	"log"
	"net/http"

	"github.com/themichaelb/Go-WebDev-Learning/internal/config"
	"github.com/themichaelb/Go-WebDev-Learning/internal/routes"
)

func main() {
	config.LoadConfig()

	router := routes.SetupRoutes()

	// Serve static files
	fs := http.FileServer(http.Dir("web/static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
