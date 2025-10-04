package main

import (
	"fmt"
	// "final-project-golang-bootcamp/connection"
	"final-project-golang-bootcamp/routers"
)

func main() {
	var PORT = ":8080"
	fmt.Println("Server will start at" + PORT)
	server := routers.StartServer()
	server.Run(PORT)
}