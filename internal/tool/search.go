package tool

import "context"

type SearchTool struct{}

func (s SearchTool) Name() string {
	return "search"
}

func (s SearchTool) Description() string {
	return "search query from user"
}

func (s SearchTool) Call(ctx context.Context, input string) (string, error) {
	return "search on llm", nil
}

