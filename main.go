package at

import (
	"github.com/MikeMwita/at/domain/services"
	"log"
	"net/http"
)

func main() {
	// Load the configuration
	//cfg, err := config.LoadConfig()
	//if err != nil {
	//	log.Fatal(err)
	//}

	// Create the SMS service
	smsService := services.NewSMSService()

	//// Create the SMS repository
	//smsRepo, err := adapters.NewMySQLSMSAdapter(cfg.MySQL)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// Create the SMS handler
	smsHandler := adapters.NewHTTPSMSAdapter(smsService, smsRepo)

	// Register the HTTP handler
	http.HandleFunc("/sendSMS", smsHandler.SendSMS)

	// Start the HTTP server
	log.Println("Starting the server on port", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
