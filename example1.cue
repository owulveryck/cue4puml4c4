package main

import (
	"github.com/owulveryck/cue4puml4c4/technology/gcp"
	"github.com/owulveryck/cue4puml4c4/technology/dev"
	"github.com/owulveryck/cue4puml4c4:c4"
	"list"
)

// Tags
elementsTags: [ID=_]: c4.#ElementTag & {
	id: "\(ID)"
}

relationsTags: [ID=_]: c4.#RelationTag & {
	id: "\(ID)"
}
elementsTags: aSupprimer: {
	legendText:  "A Migrer en React"
	bgColor:     "#f0239c"
	borderColor: "#ff0239c"
	technology:  dev.React
	shadowing:   true
	shape:       "eightsided"
}
relationsTags: myTest: {lineStyle: "bold"}
relationsTags: autreFleche: {lineColor: "#ff0000"}

// Containers

myWebApp: c4.#Container & {
	id:          "web_app"
	label:       "Web Application"
	description: "Allows users to compare multiple Twitter timelines"
	technology:  gcp.CloudStorage
	tags: [elementsTags.aSupprimer]
}

myothercontainer: c4.#Container & {
	id:         "othercontainer"
	label:      "Cool"
	technology: dev.React
}

othersample: c4.#System & {
	id:         "sample2"
	label:      "Twitter"
	technology: gcp.CloudStorage
	link:       "https://www.twitter.com"
}

twitter: c4.#System & {
	id:         "twitter"
	label:      "Twitter"
	technology: gcp.Vertexai
	link:       "https://www.twitter.com"
}

sampleSystem: c4.#System & {
	id:    "c1"
	label: "Sample System"
	containers: [myWebApp, myothercontainer]
	technology: gcp.CloudSql
	systems: [twitter, othersample]
}

admin: c4.#Person & {
	id:    "admin"
	label: "Administrator"
}

C1: c4.#C1 & {
	Persons: [admin]
	Technologies: [ dev.React, dev.CSharp, gcp.CloudStorage]
	Systems: [sampleSystem]
	Relations: [
		{source: admin, dest:    myWebApp, description: "Uses", protocol:            "HTTPS ", tags: [relationsTags.myTest]},
		{source: myWebApp, dest: twitter, description:  "Get tweets from", protocol: "HTTPS ", link: "https://plantuml.com/link"},
	]
}

#RecurseN: {
	// this is the bound on our recursion
	#maxiter: uint | *4

	// This is the function list element
	// we generate this to simulate recursion
	#funcFactory: {
		#next: _
		#func: _
	}

	// this is our "recursion unrolling"
	for k, v in list.Range(0, #maxiter, 1) {
		// this is where we build up our indexed functions and the references between them
		#funcs: "\(k)": (#funcFactory & {#next: #funcs["\(k+1)"]}).#func
	}

	// our final value needs to be null
	#funcs: "\(#maxiter)": null

	// we embed the head of the list so that the values
	// we write this into can be used like other CUE "functions"
	#funcs["0"]
}
#GetAllTech: {
	#next: _
	#func: {
		#in: _
		tech: {
			for i, x in #in {
				if x.technology != _|_ {
					"\(x.technology.name)": x.technology
				}
				if x.systems != _|_ {
					(#next & {#in: x.systems}).tech
				}
				if x.containers != _|_ {
					(#next & {#in: x.containers}).tech
				}
			}
		}
	}
}
#Depth: #RecurseN & {#funcFactory: #GetAllTech}

techs: #Depth & {#in: C1.Systems}
