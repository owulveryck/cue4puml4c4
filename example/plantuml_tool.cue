package main

import (
	"tool/cli"
	"text/template"
	"tool/file"
)

command: genpuml: {
	c1: cli.Print & {
		text: template.Execute(_plantUMLC4Template.contents, C1)
	}
}
_plantUMLC4Template: file.Read & {
	filename: "plantuml_c4.tmpl"
	contents: string
}
