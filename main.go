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


	if os.Getenv("PORT") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
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
