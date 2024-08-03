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
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	app := router.Router()

	

	if PORT== ""{
		PORT = "4000"
	}

	


	log.Fatal(app.Listen("0.0.0.0:"+PORT))

	
	
}