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
