package tool

import "context"

type WeatherTool struct{}

func (w WeatherTool) Name() string {
	return "weather"
}

func (w WeatherTool) Description() string {
	return "Get weather for a city"
}

func (w WeatherTool) Call(ctx context.Context, input string) (string, error) {

	return "give me weather of " + input, nil
}
