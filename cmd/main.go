package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ShreyasNarvekar/langChain/internal/agent"
	"github.com/ShreyasNarvekar/langChain/internal/llm"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// information := "Elon Musk"

	// summaryTemplate := "give me some information about:" + information + "\n Summary:\n1. 2 interesting facts\n2. latest project about him/her"

	//Getting the instance of ollama llm with respective model
	fmt.Println("provider:::", os.Getenv("provider"))
	llm, err := llm.NewLLM(os.Getenv("provider"))
	if err != nil {
		log.Fatalf("Failed to initialize Ollama: %v", err)
	}
	fmt.Println("2")
	executor := agent.GetNewAgent(llm)
	fmt.Println("3")

	resp, err := agent.RunAgent(executor, "London")
	fmt.Println("4")

	// resp, err := llms.GenerateFromSinglePrompt(context.Background(), llm, summaryTemplate)
	if err != nil {
		log.Fatalf("Failed to generate respone from prompt: %v", err)
	}

	fmt.Println("Response:\n", resp["output"])

}
