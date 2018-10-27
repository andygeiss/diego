package assert

import (
	"testing"
)

// That ...
func That(t *testing.T, state interface{}, m Matcher) {
	if !m.Matches(state) {
		t.Errorf("ERROR: [%v] does not match!", state)
	}
}
