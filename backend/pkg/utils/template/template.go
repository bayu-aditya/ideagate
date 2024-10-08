package template

import (
	"bytes"
	"strings"
	"text/template"
)

func GetValue(data any, query string) string {
	tmpl, err := template.New("").Parse(query)

	if err != nil {
		return ""
	}

	var resultBuffer bytes.Buffer
	if err = tmpl.Execute(&resultBuffer, data); err != nil {
		return ""
	}

	result := resultBuffer.String()
	result = strings.ReplaceAll(result, "<no value>", "")

	return result
}
