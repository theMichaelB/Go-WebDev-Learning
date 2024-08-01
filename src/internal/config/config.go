package config

import (
	"log"
	"os"
)

var (
	// AppName holds the name of the application.
	// It is used for logging and display purposes.
	AppName string
	// Port holds the port the application will listen on.
	Port    string
)

func LoadConfig() {
	AppName = os.Getenv("APP_NAME")	
	if AppName == "" {
		AppName = "Go-WebDev-Learning"
	}

	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
	}

	log.Printf("Config loaded: AppName=%s, Port=%s", AppName, Port)
}
