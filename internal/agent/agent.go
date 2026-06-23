package agent

import (
	"context"
	"fmt"

	"github.com/ShreyasNarvekar/langChain/internal/tool"
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/tools"
)

func GetNewAgent(llm llms.Model) *agents.Executor {
	calculator := tool.CalculatorTool{}
	weather := tool.WeatherTool{}
	search := tool.SearchTool{}

	executor := agents.NewExecutor(
		agents.NewOneShotAgent(
			llm,
			[]tools.Tool{
				calculator,
				search,
				weather,
			},
		),
	)

	return executor
}

func RunAgent(executor *agents.Executor, question string) (map[string]any, error) {

	result, err := executor.Call(
		context.Background(),
		map[string]any{
			"input": question,
		},
	)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("123456")
	return result, err

}
