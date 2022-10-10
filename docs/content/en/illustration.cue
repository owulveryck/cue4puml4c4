package main

import (
	"github.com/owulveryck/cue4puml4c4:c4"
	"github.com/owulveryck/cue4puml4c4/technology/stdlib"
	"github.com/owulveryck/cue4puml4c4/technology/dev"
)

user: c4.#Person & {
	id:    "User"
	label: "You"
}

cue4puml4c4: c4.#System & {
	id:         "cue4puml4c4"
	isBoundary: true
	containers: [yourCode, genpuml, yourPuml]
	relations: [
		{source: yourCode, dest: genpuml, description:  "Apply plantuml template", tags: [internalWiring]},
		{source: genpuml, dest:  yourPuml, description: "Dump the plantuml file", tags: [internalWiring]},
	]
}

yourCode: c4.#Container & {
	id:          "data"
	label:       "C4 Described with CUE"
	technology:  dev.CUE
	description: "Describe your diagram as data with the library"
	tags: [userT]
}

genpuml: c4.#Container & {
	id:          "genpuml"
	label:       "cue genpuml"
	description: "run the cue genpuml command to generate the representation"
	technology:  dev.CUE
	tags: [commandT]
}

yourPuml: c4.#Container & {
	id:          "yourPuml"
	label:       "PUML representation of the C4"
	description: "Your c4 is represented as plantuml code"
	technology:  plantumlTech
	tags: [plantumlT]
}

plantuml: c4.#System & {
	id:    "plantuml"
	label: "Plantuml"
	containers: [plantumlServer]
	tags: [plantumlT]
}

plantumlServer: c4.#Container & {
	id:         "plantumlServer", label: "Plantuml Server"
	technology: plantumlTech
}

result: c4.#System & {
	id:         "result"
	isBoundary: true
	containers: [svg, png]
}

svg: c4.#Container & {
	id:         "pictureSvg"
	label:      "SVG"
	technology: stdlib.Svg
}
png: c4.#Container & {
	id:    "picturePng"
	label: "png"
}

C1: c4.#C1 & {// the name C1 should be coherent with the name you declare in the command
	layout: "top-down"
	Systems: [cue4puml4c4, plantuml, result]
	Persons: [user]
	relations: [
		{source: user, dest:     yourCode, description: "Write the diagram", tags: [userInput]},
		{source: user, dest:     genpuml, description:  "Run the command", tags: [userInput]},
		{source: yourPuml, dest: plantumlServer, tags: [externalSystem]},
		{source: plantuml, dest: result, tags: [externalSystem]},
	]
}

userInput: c4.#RelationTag & {
	id:         "userInput"
	legendText: "User Input"
	lineStyle:  "bold"
	textColor:  "#008800"
	lineColor:  "#00bb00"
}

externalSystem: c4.#RelationTag & {
	id:         "externalSystem"
	legendText: "external component"
	lineStyle:  "dash"
	textColor:  "#ffffff"
	lineColor:  "#ffffff"
}

plantumlT: c4.#ElementTag & {
	id:          "plantumlT"
	legendText:  "PlantUML"
	shape:       "rounded"
	bgColor:     "#004455"
	borderColor: "#aabbcc"
}
userT: c4.#ElementTag & {
	id:          "User Generated Data"
	legendText:  ""
	shape:       "rounded"
	bgColor:     "#008800"
	borderColor: "#00bb00"
}
commandT: c4.#ElementTag & {
	id:          "commandT"
	legendText:  "Command"
	shape:       "rounded"
	bgColor:     "#004400"
	borderColor: "#00bb00"
}
internalWiring: c4.#RelationTag & {
	id:         "internalWiring"
	legendText: "internal machinery"
	lineStyle:  "dash"
	textColor:  "#ffffff"
	lineColor:  "#ffffff"
}

plantumlTech: c4.#Technology & {
	name: "plantuml"
}
