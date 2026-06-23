package tool

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

type CalculatorTool struct{}

func (c CalculatorTool) Name() string {
	return "calculator"
}

func (c CalculatorTool) Description() string {
	return "Calculate the query which user gave"
}

func (c CalculatorTool) Call(ctx context.Context, input string) (string, error) {
	parts := strings.Split(input, ",")

	if len(parts) != 2 {
		return "", fmt.Errorf("invalid input")
	}

	a, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
	b, _ := strconv.Atoi(strings.TrimSpace(parts[1]))

	return fmt.Sprintf("%d", a*b), nil
}
