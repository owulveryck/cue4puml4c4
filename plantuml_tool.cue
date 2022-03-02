package c4

import (
	"tool/cli"
	"text/template"
)

command: genpuml: {
	c1: cli.Print & {
		text: template.Execute(_systemTemplate, C1)
	}
}

_systemTemplate: """
	@startuml MyName
	!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml
	
	!include https://raw.githubusercontent.com/owulveryck/PlantUML-icons-GCP/master/official/GCPCommon.puml
	{{- range .Technologies}}
	{{- if .sprite }}
	!include {{ .sprite.url }} 
	{{end -}}
	{{end -}}
	
	{{if .Systems}}
	{{template "Systems" .Systems}}
	{{end}}

	
	{{- define "Systems"}}
	{{- range .}}
	System{{if .isBoundary}}_Boundary{{end}}({{.id}},"{{.label}}"{{if .technology}},"{{.technology.name}}","{{.technology.sprite.id}}"{{end}}){{if or .containers .systems}}{
	{{- range .containers -}}	
	{{- template "Container" . -}}
	{{- end -}}
	{{- if .systems}}
	{{template "Systems" .systems}}
	{{end -}}
	{{- range .rels}}
	{{template "Rel" .}}
	{{end}}
	}
	{{end -}}
	{{end -}}
	{{end -}}

	{{define "Container" }}
		Container{{.technology.type}}({{.id}},"{{.label}}","{{.technology.name}}","{{if .description}}{{.description}}{{end}}"{{if .technology.sprite}},"{{.technology.sprite.id}}"{{end}})
	{{- end }}

	{{define "Rel"}}
		Rel("{{.source.id}}","{{.dest.id}}","{{.description}}")
	{{- end}}

	SHOW_LEGEND()
	@enduml
	"""
