package main

import (
	"log"
	"os"
	"final-project-golang-bootcamp/routers"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		// Log a warning if .env file is missing, but continue using system environment variables
		log.Println("Warning: .env file not found, using environment variables instead.")
	}

	// Retrieve the server port from environment variables
	var PORT = ":" + os.Getenv("PORT")

	// Log the port information before starting the server
	log.Println("Server will start at " + PORT)

	// Initialize the router and start the HTTP server
	server := routers.StartServer()
	server.Run(PORT)
}