package render

import "html/template"

var funcs = template.FuncMap{
	"add":        add,
	"isEmpty":    isEmpty,
	"isNotEmpty": isNotEmpty,
	"raw":        raw,
	"mul":        mul,
	"sub":        sub,
}

func add(a int, b int) int {
	return a + b
}

func isEmpty(slice []string) bool {
	return len(slice) == 0
}

func isNotEmpty(slice []string) bool {
	return len(slice) > 0
}

func mul(a int, b int) int {
	return a * b
}

func raw(text string) template.HTML {
	return template.HTML(text)
}

func sub(a int, b int) int {
	return a - b
}
