package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ShreyasNarvekar/langChain/internal/llm"
	_ "github.com/joho/godotenv/autoload"
	"github.com/tmc/langchaingo/llms"
)

func main() {
	information := "Elon Musk"

	summaryTemplate := "give me some information about:" + information + "\n Summary:\n1. 2 interesting facts\n2. latest project about him/her"

	//Getting the instance of ollama llm with respective model
	llm, err := llm.NewLLM(os.Getenv("provider"))
	if err != nil {
		log.Fatal("Failed to initialize Ollama: %v", err)
	}

	resp, err := llms.GenerateFromSinglePrompt(context.Background(), llm, summaryTemplate)
	if err != nil {
		log.Fatal("Failed to generate respone from prompt: %v", err)
	}

	fmt.Println("Response:\n", resp)

}
