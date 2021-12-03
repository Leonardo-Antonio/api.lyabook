package tmpl

import (
	"bytes"
	"text/template"
)

func Read(name string, data interface{}) (string, error) {
	tpl := template.Must(template.ParseGlob("template/*.tpl"))
	template.Must(tpl.ParseGlob("template/_layouts/*.tpl"))

	var bf bytes.Buffer
	if err := tpl.ExecuteTemplate(&bf, name, &data); err != nil {
		return "", err
	}
	return bf.String(), nil
}

func Report(name string, data interface{}) (string, error) {
	tpl, err := template.ParseGlob("template/report/*.tpl")
	if err != nil {
		return "", err
	}

	var bf bytes.Buffer
	if err := tpl.ExecuteTemplate(&bf, name, &data); err != nil {
		return "", err
	}
	return bf.String(), nil
}
