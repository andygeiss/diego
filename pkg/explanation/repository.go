package explanation

// Option ...
type Option struct {
	Name string
	Value interface{}
}

// Question ...
type Question struct {
	Prompt  string
	Options []*Option
}

// Repository ...
type Repository interface {
	FindQuestionsBySurvey(name string) ([]*Question, error)
}

// Survey ...
type Survey struct {
	Name string
	Questions []*Question
}
