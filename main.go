package main

import (
	"auth/server"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	err := godotenv.Load("configs/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Server is running on the port " + os.Getenv("PORT"))
	server.Start()
}
