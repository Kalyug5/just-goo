package api

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

type Content struct{
	Parts []string `json:"parts"`

}

type Candidates struct{
	Content *Content `json:"content"`
}

type ContentResponse struct{
	Candidates *[]Candidates `json:"candidates"`
}

func Api() *genai.GenerativeModel {
	err := godotenv.Load(".env")
	if err != nil {
		panic("here error happened")
	}
	ctx := context.Background()

	client,error := genai.NewClient(ctx,option.WithAPIKey(os.Getenv("API_KEY")))

	if error != nil {
		log.Fatal(error)
	}



	model := client.GenerativeModel("gemini-1.5-flash")

	return model

}