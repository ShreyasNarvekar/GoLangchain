package tool

type Tool interface {
	Name() string
	Description() string
	Call(input string) (map[string]any, error)
}
