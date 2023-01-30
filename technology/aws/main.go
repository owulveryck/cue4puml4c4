package main

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const tmpl = `
{{.Name}}: c4.#Technology & {
        name: "{{.Label}}"
        {{if eq .Group "Database"}}type: "Db"{{end}}
        sprite: {
                url: "{{.URL}}"
                id:  "{{.ID}}"
        }
}
`

func main() {
	path := os.Args[1]
	t := template.Must(template.New("cueAWS").Parse(tmpl))
	baseURL := "https://raw.githubusercontent.com/awslabs/aws-icons-for-plantuml/v14.0/dist/"
	elements := strings.Split(path, "/")
	fileName := elements[len(elements)-1]
	name := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	t.Execute(os.Stdout, struct {
		Name  string
		Label string
		URL   string
		Group string
		ID    string
	}{
		Name:  name,
		Label: "AWS " + elements[len(elements)-2] + " " + name,
		Group: elements[len(elements)-2],
		URL:   baseURL + elements[len(elements)-2] + "/" + fileName,
		ID:    name,
	})
}
