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
		{source: genpuml, dest: yourCode, description: "Read your code and apply puml template"},
		{source: genpuml, dest: yourPuml, description: "Generate puml file"},
	]
}

yourCode: c4.#Container & {
	id:         "data"
	label:      "C4 Described with CUE"
	technology: dev.CUE
}

genpuml: c4.#Container & {
	id:         "genpuml"
	label:      "Cue command genpuml"
	technology: dev.CUE
}

yourPuml: c4.#Container & {
	id:    "yourPuml"
	label: "PUML representation of the C4"
}

plantuml: c4.#System & {
	id:    "plantuml"
	label: "Plantuml"
	containers: [plantumlServer]
}

plantumlServer: c4.#Container & {
	id: "plantumlServer", label: "Plantuml Server"
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
	Systems: [cue4puml4c4, plantuml, result]
	Persons: [user]
	relations: [
		{source: user, dest:     yourCode},
		{source: yourPuml, dest: plantumlServer},
		{source: plantuml, dest: result},
	]
}
