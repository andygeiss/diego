package inference

// Engine ...
type Engine interface {
	Run(conditions []string) ([]string, error)
}
