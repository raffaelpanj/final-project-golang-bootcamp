package main

import (
	"log"
	"os"
	"final-project-golang-bootcamp/routers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables instead.")
	}
	var PORT = ":" + os.Getenv("PORT")
	log.Println("Server will start at " + PORT)
	server := routers.StartServer()
	server.Run(PORT)
}