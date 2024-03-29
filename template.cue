package c4

plantumlTemplate: """
	@startuml {{if .title}}{{.title}}{{else}}MyDiagram{{end}}
	
	!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml
	
	!include https://raw.githubusercontent.com/owulveryck/PlantUML-icons-GCP/master/official/GCPCommon.puml
	!include https://raw.githubusercontent.com/awslabs/aws-icons-for-plantuml/v14.0/dist/AWSCommon.puml
	{{- range .Technologies.tech}}
	{{- if .sprite }}
	!include {{ .sprite.url }} 
	{{- end }}
	{{- end }}
	
	{{- if eq .layout "top-down"}}
	LAYOUT_TOP_DOWN()
	{{ else if eq .layout "left-right"}}
	LAYOUT_LEFT_RIGHT()
	{{ else if eq .layout "landscape"}}
	LAYOUT_LANDSCAPE()
	{{- end}}
	{{- if .hideStereotype }}
	HIDE_STEREOTYPE()
	{{- end}}
	
	/'Relation Tags'/ 
	{{- if .relationTags}}{{template "RelationTags" .relationTags}}{{- end}}
	
	/'Element Tags'/ 
	{{- if .elementTags}}{{template "ElementTags" .elementTags}}{{- end}}

	/'People'/ 
	{{- if .Persons}}{{- template "Persons" .Persons -}} {{- end}}

	/'Systems'/ 
	{{- if .Systems}} {{template "Systems" .Systems}} {{end}}

	/'SystemsExt'/ 
	{{- if .SystemsExt}}{{template "SystemsExt" .SystemsExt}} {{end}}

	/'Relations'/ 
	{{- if .relations}}
	{{- range .relations}}
	{{template "Rel" . -}}
	{{- end}}
	{{- end}}

	SHOW_LEGEND()
	@enduml
	""" + containerTemplatePuml + relTemplatePuml + personsTemplatePuml + systemsTemplatePuml + tagsTemplatePuml + systemsExtTemplatePuml + relationTagsTemplatePuml + elementTagsTemplatePuml

mermaidTemplate: """
	C4Container
	   title {{if .title}}{{.title}}{{else}}MyDiagram{{end}}
	
	{{- if .Persons}}{{- template "Persons" .Persons -}} {{- end}}
	{{- if .Systems}} {{template "Systems" .Systems}} {{end}}
	{{- if .SystemsExt}}{{template "SystemsExt" .SystemsExt}} {{end}}
	""" + containerTemplatePuml + relTemplatePuml + personsTemplatePuml + systemsTemplateMermaid + tagsTemplatePuml + systemsExtTemplatePuml

systemsExtTemplatePuml: """
	{{- define "SystemsExt"}}
	{{- range .}}
	System_Ext({{.id}},"{{.label}}"{{if .technology}},"{{.technology.name}}","{{.technology.sprite.id}}"{{end}}{{if .link}},$link="{{.link}}"{{end}}{{if .tags}},$tags="{{template "tags".tags}}"{{end}}){{if or .containers .systems}}{
	{{- range .containers }}	
	{{template "Container" . -}}
	{{- end }}
	{{- if .systems}}
	{{template "Systems" .systems -}}
	{{- end }}
	{{if .SystemsExt}}
	{{template "SystemsExt" .SystemsExt}}
	{{end}}
	{{- range .relations}}
	{{template "Rel" . -}}
	{{- end}}
	}
	{{end -}}
	{{end -}}
	{{end -}}	
	"""

tagsTemplatePuml: """
	{{define "tags"}}{{range .}}{{.id}}+{{end}} {{end}}
	"""

systemsTemplatePuml: """
	{{- define "Systems" }}
	{{- range . }}
	System{{if .isBoundary}}_Boundary{{end}}({{.id}},"{{.label}}"{{if .description}},"{{.desc}}"{{end}}{{if .technology}},{{if not .description}}"{{.technology.name}}",{{end}}"{{.technology.sprite.id}}"{{end}}{{if .link}},$link="{{.link}}"{{end}}{{if .tags}},$tags="{{template "tags".tags}}"{{end}}){{    if or .containers .systems}} {
	{{- range .containers }}	
	{{ template "Container" . -}}
	{{- end }}
	{{- if .systems }}
	{{- template "Systems" .systems -}}
	{{- end }}
	{{- range .relations }}
	{{template "Rel" . -}}
	{{- end }}
	}
	{{- end}}
	{{- end}}
	{{- end}}
	"""

systemsTemplateMermaid: """
	{{- define "Systems" }}
	{{- range . }}
	System{{if or .containers .systems }}_Boundary{{end}}({{.id}},"{{.label}}"{{if .description}},"{{.desc}}"{{end}}{{if .technology}},{{if not .description}}"{{.technology.name}}",{{end}}"{{.technology.sprite.id}}"{{end}}{{if .link}},$link="{{.link}}"{{end}}{{if .tags}},$tags="{{template "tags".tags}}"{{end}}){{    if or .containers .systems}} {
	{{- range .containers }}	
	{{ template "Container" . -}}
	{{- end }}
	{{- if .systems }}
	{{- template "Systems" .systems -}}
	{{- end }}
	{{- range .relations }}
	{{template "Rel" . -}}
	{{- end }}
	}
	{{- end}}
	{{- end}}
	{{- end}}
	"""

personsTemplatePuml: """
	{{ define "Persons"}}{{- range .}}
	Person({{.id}},"{{.label}}")
	{{- end }}{{ end }}
	"""

relTemplatePuml: """
	{{ define "Rel"}}Rel("{{.source.id}}","{{.dest.id}}","{{.desc}}"{{if .protocol}},"{{.protocol}}"{{end}}{{if .link}},$link="{{.link}}"{{end}}{{if .tags}},$tags="{{template "tags".tags}}"{{end}}){{ end}}
	"""

containerTemplatePuml: """
	{{ define "Container" }}Container{{.technology.type}}({{.id}},"{{.label}}","{{.technology.name}}","{{if .description}}{{.desc}}{{end}}"{{if .technology.sprite}},"{{.technology.sprite.id}}"{{end}}{{if .link}},$link="{{.link}}"{{end}}{{if .tags}},$tags="{{template "tags".tags}}"{{end}}) {{ end }}
	"""

relationTagsTemplatePuml: """
	{{- define "RelationTags"}}
	{{- range .a.tags}}
	AddRelTag("{{.id}}"
	{{- if .textColor}},$textColor="{{.textColor}}"{{end -}}
	{{- if .lineColor}},$lineColor="{{.lineColor}}"{{ end -}}
	{{- if .lineStyle}},$lineStyle={{if eq .lineStyle  "dot"}}DottedLine(){{end}}{{if eq .lineStyle  "dash"}}DashedLine(){{end}}{{if eq .lineStyle  "bold"}}BoldLine(){{end}}{{ end -}}
	{{- if .legendText}},$legendText="{{.legendText}}"{{ end -}}
	{{- if .technology}},$techn="{{.technology.name}}",$sprite="{{.technology.sprite.id}}"{{ end -}})
	{{- end}}
	{{- range .b.tags}}
	AddRelTag("{{.id}}"
	{{- if .textColor}},$textColor="{{.textColor}}"{{end -}}
	{{- if .lineColor}},$lineColor="{{.lineColor}}"{{ end -}}
	{{- if .lineStyle}},$lineStyle={{if eq .lineStyle  "dot"}}DottedLine(){{end}}{{if eq .lineStyle  "dash"}}DashedLine(){{end}}{{if eq .lineStyle  "bold"}}BoldLine(){{end}}{{ end -}}
	{{- if .legendText}},$legendText="{{.legendText}}"{{ end -}}
	{{- if .technology}},$techn="{{.technology.name}}",$sprite="{{.technology.sprite.id}}"{{ end -}})
	{{- end}}
	{{- end}}
	"""

elementTagsTemplatePuml: """
	{{- define "ElementTags"}}
	{{- range .tags}}
	AddElementTag("{{.id}}"
	{{- if .bgColor}},$bgColor="{{.bgColor}}"{{end -}}
	{{- if .borderColor}},$borderColor="{{.borderColor}}"{{ end -}}
	{{- if .shadowing}},$shadowing="{{.shadowing}}"{{ end -}}
	{{- if .shape}},$shape="{{if eq .shape  "rounded"}}RoundedBoxShape(){{end}}{{if eq .shape  "eightsided"}}EightSidedShape(){{end}}"{{ end -}}
	{{- if .fontColor}},$fontColor="{{.fontColor}}"{{ end -}}
	{{- if .legendText}},$legendText="{{.legendText}}"{{ end -}}
	{{- if .technology}},$techn="{{.technology.name}}",$sprite="{{.technology.sprite.id}}"{{ end -}})
	{{- end}}
	{{- end}}
	"""
