package main

import (
	"tool/cli"
	"text/template"
	// "tool/file"
	"github.com/owulveryck/cue4puml4c4:c4"
)

command: genpuml: {
	c1: cli.Print & {
		//text: template.Execute(_plantUMLC4Template.contents, C1)
		text: template.Execute(c4.plantumlTemplate, C1)
	}
}
command: genmermaid: {
	c1: cli.Print & {
		//text: template.Execute(_plantUMLC4Template.contents, C1)
		text: template.Execute(c4.mermaidTemplate, C1)
	}
}
//_plantUMLC4Template: file.Read & {
// filename: "plantuml_c4.tmpl"
// contents: string
//}
