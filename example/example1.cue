package main

import (
	"github.com/owulveryck/cue4puml4c4/technology/gcp"
	"github.com/owulveryck/cue4puml4c4/technology/dev"
	"github.com/owulveryck/cue4puml4c4:c4"
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
	Systems: [sampleSystem]
	Relations: [
		{source: admin, dest:    myWebApp, description: "Uses", protocol:            "HTTPS ", tags: [relationsTags.myTest]},
		{source: myWebApp, dest: twitter, description:  "Get tweets from", protocol: "HTTPS ", link: "https://plantuml.com/link"},
	]
}
