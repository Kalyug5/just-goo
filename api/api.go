package api

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

type Content struct {
	Parts []string `json:"parts"`
}

type Candidates struct {
	Content *Content `json:"content"`
}

type ContentResponse struct {
	Candidates *[]Candidates `json:"candidates"`
}

func Api() *genai.GenerativeModel {

	if os.Getenv("API_KEY") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	ctx := context.Background()

	client, error := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))

	if error != nil {
		log.Fatal(error)
	}

	model := client.GenerativeModel("gemini-2.5-flash")

	return model

}
