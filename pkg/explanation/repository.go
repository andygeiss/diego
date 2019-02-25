package explanation

// Option ...
type Option struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Value       interface{} `json:"value"`
}

// Question ...
type Question struct {
	Prompt      string    `json:"prompt"`
	Description string    `json:"description"`
	Options     []*Option `json:"options"`
}

// Repository ...
type Repository interface {
	FindQuestionsBySurvey(name string) ([]*Question, error)
}

// Survey ...
type Survey struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Questions   []*Question `json:"questions"`
}
