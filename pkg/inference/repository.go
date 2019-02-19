package inference

// Concept ...
type Concept struct {
	Name string
	Facts []string
	Rules []*Rule
}

// Repository ...
type Repository interface {
	FindConceptByName(name string) (*Concept, error)
}

// Rule ...
type Rule struct {
	Conditions  []string
	Conclusions []string
}
