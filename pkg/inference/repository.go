package inference

// Concept ...
type Concept struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Facts       []string `json:"facts"`
	Rules       []*Rule  `json:"rules"`
}

// Repository ...
type Repository interface {
	FindConceptByName(name string) (*Concept, error)
}

// Rule ...
type Rule struct {
	Conditions  []string `json:"conditions"`
	Conclusions []string `json:"conclusions"`
}
