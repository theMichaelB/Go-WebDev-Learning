package routes

import (
	"net/http"

	"github.com/themichaelb/Go-WebDev-Learning/internal/handlers"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HomeHandler)
	return mux
}
