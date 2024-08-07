package main

import (
	"log"
	"os"
	"text/template"
)

type Code struct {
	Title string
	Items []string
}

func main() {

	f, err := os.Create("fmt.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data := Code{"Example", []string{"Alpha", "Beta", "Gamma", "Delta"}}
	text, err := template.New("example").Parse(templates)
	if err != nil {
		log.Fatal(err)
	}
	err = text.Execute(f, data)
	if err != nil {
		log.Fatal(err)
	}
}

var templates = `

#{{.Title}}

{{range .Items }}
	+{{ . -}}
{{- end}}

	`
