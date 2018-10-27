package assert

// Matcher ...
type Matcher interface {
	Matches(interface{}) bool
	String() string
}
