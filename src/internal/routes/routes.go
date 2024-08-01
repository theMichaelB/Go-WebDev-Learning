package routes

import (
	"net/http"

	"github.com/themichaelb/Go-WebDev-Learning/internal/handlers"
)

// SetupRoutes sets up the routes for the application.
func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HomeHandler)

	return mux
}
