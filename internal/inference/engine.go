package inference

import "errors"

// Engine ...
type Engine interface {
	Run(conditions []string) ([]string, error)
}

type defaultEngine struct {
	concept string
	repo    Repository
}

// NewDefaultEngine ...
func NewDefaultEngine(concept string, repo Repository) Engine {
	return &defaultEngine{
		concept: concept,
		repo:    repo,
	}
}

// Run ...
func (e *defaultEngine) Run(conditions []string) ([]string, error) {

	concept, err := e.repo.FindConceptByName(e.concept)
	if err != nil {
		return nil, err
	}

	if concept == nil {
		return nil, errors.New("concept was not found")
	}

	facts := concept.Facts
	rules := concept.Rules

	// Add the current conditions as facts
	facts = append(facts, conditions...)

	applied := make(map[string]bool, 0)

	for i := 0; i < len(rules); i++ {
		for _, rule := range rules {
			// Match ...
			if e.match(rule, facts) {
				// Handle collision resolution ...
				rule = e.recentlyEntered(rule, rules)
				// Fire ...
				for _, action := range rule.Conclusions {
					if _, ok := applied[action]; !ok {
						facts = append(facts, action)
						applied[action] = true
					}
				}
			}
		}
	}
	return facts, nil
}

func (e *defaultEngine) recentlyEntered(rule *Rule, rules []*Rule) *Rule {
	var result *Rule
	for _, r := range rules {
		// Match conditions ...
		count := len(r.Conditions)
		for _, cond := range r.Conditions {
			for _, c := range rule.Conditions {
				if c == cond {
					count--
				}
			}
		}
		// Set the recently entered rule.
		if count == 0 {
			result = r
		}
	}
	return result
}

func (e *defaultEngine) match(rule *Rule, facts []string) bool {
	wanted := len(rule.Conditions)
	for _, cond := range rule.Conditions {
		for _, fact := range facts {
			if cond == fact {
				wanted--
			}
		}
	}
	return wanted == 0
}
