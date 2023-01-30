package main

import (
	"github.com/owulveryck/cue4puml4c4/technology/gcp"
	"github.com/owulveryck/cue4puml4c4/technology/aws"
	"github.com/owulveryck/cue4puml4c4/technology/dev"
	"github.com/owulveryck/cue4puml4c4/technology/stdlib"
	"github.com/owulveryck/cue4puml4c4:c4"
)

aSupprimer: c4.#ElementTag & {
	id:          "aSupprimer"
	legendText:  "A Migrer en Rust"
	bgColor:     "#f0235c"
	borderColor: "#ff0239c"
	technology:  dev.Rust
	shadowing:   true
	shape:       "eightsided"
}
myTest:      c4.#RelationTag & {id: "myTest", lineStyle:      "bold"}
autreFleche: c4.#RelationTag & {id: "autreFleche", lineColor: "#ff0000"}

// Containers

appEngine: c4.#System & {
	id:         "appEngine"
	label:      "App Engine"
	technology: gcp.AppEngine
	containers: [myWebApp]
}

myWebApp: c4.#Container & {
	id:    "web_appp"
	label: "Web Application"
	description: """
		Allows users
		to compare multiple Twitter timelines
		"""
	technology: dev.Go
	tags: [aSupprimer]
}

mysample2: c4.#Container & {
	id:          "mysample2"
	label:       "mysample2"
	description: "mysample2"
	technology:  aws.EC2
}
myothercontainer: c4.#Container & {
	id:         "othercontainer"
	label:      "Cool2"
	technology: stdlib.Kafka
}

othersample: c4.#Container & {
	id:         "sample2"
	label:      "Event Source"
	technology: gcp.Pubsub
	link:       "https://www.twitter.com"
}

twitter: c4.#System & {
	id:         "twitter"
	label:      "Twitter"
	technology: gcp.ComputeEngine
	link:       "https://www.twitter.com"
	containers: [{
		id:         "myDatabase"
		technology: dev.Postgresql
		label:      "My Awesome DB"
	}]
}

sampleSystem: c4.#System & {
	id:         "c1"
	label:      "Sample System"
	isBoundary: true
	containers: [othersample, myothercontainer, mysample2]
	systems: [appEngine, twitter]
	technology: dev.CUE
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
	layout: "landscape"
	relations: [
		{source: admin, dest:    myWebApp, description: "Uses", protocol:            "HTTPS ", tags: [myTest]},
		{source: myWebApp, dest: twitter, description:  "Get tweets from", protocol: "HTTPS ", link: "https://plantuml.com/link"},
	]
}
