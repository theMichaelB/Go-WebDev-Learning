package config

import (
	"log"
	"os"
)

var (
	AppName string
	Port    string
)

func LoadConfig() {
	AppName = os.Getenv("APP_NAME")
	Port = os.Getenv("PORT")

	if AppName == "" {
		AppName = "Go-WebDev-Learning"
	}
	if Port == "" {
		Port = "8080"
	}

	log.Printf("Config loaded: AppName=%s, Port=%s", AppName, Port)
}
