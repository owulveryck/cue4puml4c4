package main

import (
	"github.com/owulveryck/cue4puml4c4/technology/gcp"
	"github.com/owulveryck/cue4puml4c4/technology/dev"
	"github.com/owulveryck/cue4puml4c4/technology/stdlib"
	"github.com/owulveryck/cue4puml4c4:c4"
)

aSupprimer: c4.#ElementTag & {
	id:          "aSupprimer"
	legendText:  "A Migrer en React"
	bgColor:     "#f0239c"
	borderColor: "#ff0239c"
	technology:  dev.React
	shadowing:   true
	shape:       "eightsided"
}
myTest: {id: "myTest", lineStyle: "bold"}
autreFleche: {id: "autreFleche", lineColor: "#ff0000"}

// Containers

myWebApp: c4.#Container & {
	id:          "web_appp"
	label:       "Web Application"
	description: "Allows users to compare multiple Twitter timelines"
	technology:  dev.CUE
	tags: [aSupprimer]
}

myothercontainer: c4.#Container & {
	id:         "othercontainer"
	label:      "Cool2"
	technology: stdlib.Kafka
}

othersample: c4.#Container & {
	id:         "sample2"
	label:      "Twitter2"
	technology: gcp.Pubsub
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
	containers: [myWebApp, othersample, myothercontainer]
	technology: gcp.CloudSql
	systems: [twitter]
	relations: [
		{source: othersample, dest: myothercontainer, tags: [autreFleche]},
	]
}

admin: c4.#Person & {
	id:    "admin"
	label: "Administrator"
}

C1: c4.#C1 & {
	Persons: [admin]
	Systems: [sampleSystem]
	relations: [
		{source: admin, dest:    myWebApp, description: "Uses", protocol:            "HTTPS ", tags: [myTest]},
		{source: myWebApp, dest: twitter, description:  "Get tweets from", protocol: "HTTPS ", link: "https://plantuml.com/link"},
	]
}
