package main

import (
	"fmt"
	"github.com/MikeMwita/africastalking-go/config"

	"log"
	"net/http"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Load configuration
	cfg := config.LoadConfig()

	// Create the SMS service
	smsService := services.NewSMSService(cfg.Username, cfg.APIKey, cfg.Env)
	fmt.Println(cfg.Username, cfg.APIKey, cfg.Env)

	// Create the SMS handler
	smsHandler := adapters.NewHTTPSMSHandler(smsService)

	// Register the HTTP handler
	http.HandleFunc("/sendSMS", smsHandler.ServeHTTP)

	// Start the HTTP server
	addr := ":" + cfg.Port
	log.Println("Starting the server on address", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
