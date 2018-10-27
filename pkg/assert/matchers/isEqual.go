package matchers

import (
	"fmt"
	"github.com/andygeiss/diego/pkg/assert"
	"reflect"
)

type isEqualMatcher struct {
	val interface{}
}

// Matches ...
func (m *isEqualMatcher) Matches(val interface{}) bool {
	return reflect.DeepEqual(m.val, val)
}

// String ...
func (m *isEqualMatcher) String() string {
	return fmt.Sprintf("[%s] is equal", m.val)
}

// IsEqual ...
func IsEqual(val interface{}) assert.Matcher {
	return &isEqualMatcher{val: val}
}
