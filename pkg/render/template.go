package render

import (
	"bytes"
	"html/template"
	"path"
)

// Template ...
func Template(filename string, data interface{}) ([]byte, error) {
	var out bytes.Buffer
	tmpl, err := template.New(path.Base(filename)).Funcs(funcs).ParseFiles(filename)
	if err != nil {
		return nil, err
	}
	err = tmpl.Execute(&out, data)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
