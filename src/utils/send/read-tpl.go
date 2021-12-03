package send

import (
	"bytes"
	"html/template"
)

func readTemplate(name, path string, data interface{}) (string, error) {
	tpl, err := template.New(name).ParseFiles(path)
	if err != nil {
		return "", err
	}

	var t bytes.Buffer
	if err := tpl.Execute(&t, &data); err != nil {
		return "", err
	}

	return t.String(), nil
}
