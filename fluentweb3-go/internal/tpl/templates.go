package tpl

import (
	"log"
	"strings"
	"text/template"
)

func RenderTemplate(tplString string, data any) string {
	tpl, err := template.New("tpl").Parse(tplString)
	if err != nil {
		log.Fatal(err)
	}
	var outputString strings.Builder
	err = tpl.Execute(&outputString, data)
	if err != nil {
		log.Fatal(err)
	}
	return outputString.String()
}
