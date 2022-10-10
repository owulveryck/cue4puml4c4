---
title: About CUE4Puml4C4
linkTitle: About
type: docs
menu:
  main:
    weight: 5

---

## Example

Here is an example of the potential of this library

{{< tabpane >}}
{{< tab header="Diagram data (CUE)" lang="go" >}}
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
	layout: "landscape"
	Systems: [cue4puml4c4, plantuml, result]
	Persons: [user]
	relations: [
		{source: user, dest:     yourCode},
		{source: yourPuml, dest: plantumlServer},
		{source: plantuml, dest: result},
	]
}
{{< /tab >}}
{{< tab header="CUE command" lang="go" >}}
package main

import (
        "tool/cli"
        "text/template"
        "github.com/owulveryck/cue4puml4c4:c4"
)

command: genpuml: {
        c1: cli.Print & {
                text: template.Execute(c4.plantumlTemplate, C1) // change C1 here with the name of your object
        }
}
{{< /tab >}}
{{< tab header="Plantuml code" >}}
@startuml MyDiagram
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml
!include https://raw.githubusercontent.com/owulveryck/PlantUML-icons-GCP/master/official/GCPCommon.puml
!include https://raw.githubusercontent.com/owulveryck/cue4puml4c4/main/assets/cue.puml 
!include https://raw.githubusercontent.com/plantuml-stdlib/gilbarbara-plantuml-sprites/master/sprites/./svg.puml 
LAYOUT_LANDSCAPE()
HIDE_STEREOTYPE()
/'Relation Tags'/ 
/'Element Tags'/ 
/' Persons '/
Person(User,"You")
	
/' Systems '/
System_Boundary(cue4puml4c4,"cue4puml4c4"){
	Container(data,"C4 Described with CUE","CUE","","cue_logo")
	Container(genpuml,"Cue command genpuml","CUE","","cue_logo")
	Container(yourPuml,"PUML representation of the C4","Undefined","")
	Rel("genpuml","data","Read your code and apply puml template")
	Rel("genpuml","yourPuml","Generate puml file")
}
System(plantuml,"Plantuml"){
	Container(plantumlServer,"Plantuml Server","Undefined","")
}
System_Boundary(result,"result"){
	Container(pictureSvg,"SVG","Svg","","svg")
	Container(picturePng,"png","Undefined","")
}
/' Relations '/
	Rel("User","data","")
	Rel("yourPuml","plantumlServer","")
	Rel("plantuml","result","")
SHOW_LEGEND()
@enduml
{{< /tab >}}
{{< tab header="Plantuml rendering" lang="plantuml" >}}
@startuml MyDiagram
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml
!include https://raw.githubusercontent.com/owulveryck/PlantUML-icons-GCP/master/official/GCPCommon.puml
!include https://raw.githubusercontent.com/owulveryck/cue4puml4c4/main/assets/cue.puml 
!include https://raw.githubusercontent.com/plantuml-stdlib/gilbarbara-plantuml-sprites/master/sprites/./svg.puml 
LAYOUT_LANDSCAPE()
HIDE_STEREOTYPE()
/'Relation Tags'/ 
/'Element Tags'/ 
/' Persons '/
Person(User,"You")
	
/' Systems '/
System_Boundary(cue4puml4c4,"cue4puml4c4"){
	Container(data,"C4 Described with CUE","CUE","","cue_logo")
	Container(genpuml,"Cue command genpuml","CUE","","cue_logo")
	Container(yourPuml,"PUML representation of the C4","Undefined","")
	Rel("genpuml","data","Read your code and apply puml template")
	Rel("genpuml","yourPuml","Generate puml file")
}
System(plantuml,"Plantuml"){
	Container(plantumlServer,"Plantuml Server","Undefined","")
}
System_Boundary(result,"result"){
	Container(pictureSvg,"SVG","Svg","","svg")
	Container(picturePng,"png","Undefined","")
}
/' Relations '/
	Rel("User","data","")
	Rel("yourPuml","plantumlServer","")
	Rel("plantuml","result","")
SHOW_LEGEND()
@enduml
{{< /tab >}}

{{< /tabpane >}}

This website is a work-in-progress. Its goal is to help the user use the library.
The layout of the website follows the [divio documentation system](https://documentation.divio.com/). You will find four sections:

## Tutorials

> Tutorials are lessons that take the reader by the hand through a series of steps to complete a project of some kind. They are what your project needs in order to show a beginner that they can achieve something with it.

## How-to guides

> How-to guides take the reader through the steps required to solve a real-world problem.

## Reference

> Reference guides are technical descriptions of the machinery and how to operate it.

## Explanation

> Explanation, or discussions, clarify and illuminate a particular topic. They broaden the documentation’s coverage of a topic.


