package email

import (
	"bytes"
	"text/template"
)

var dirTemplate string = "../../templates/"

func TemplateFileName(body *bytes.Buffer, fileName string) error {
	dirTemplate += fileName
	var templates = template.Must(template.ParseFiles(dirTemplate))
	if err := templates.ExecuteTemplate(body, fileName, nil); err != nil {
		return err
	}

	return nil
}
