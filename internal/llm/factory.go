package llm

import (
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/anthropic"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/llms/openai"
)

func NewLLM(provider string) (llms.Model, error) {
	switch provider {
	case "ollama":
		return ollama.New(ollama.WithModel("gemma3:270m"))

	case "openai":
		return openai.New(openai.WithToken("OPENAI_API_KEY"))

	case "anthropic":
		return anthropic.New(anthropic.WithToken("OPENAI_API_KEY"))

	default:
		return nil, fmt.Errorf("unsupported provider: %s enter in lowercase", provider)
	}
}
