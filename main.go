package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Kalyug5/just-goo/router"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Go Lang Project")


	if os.Getenv("GO_ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "4000" 
	}

	app := router.Router()

	log.Printf("Listening on port %s", PORT)
	log.Fatal(app.Listen("0.0.0.0:" + PORT))
}
