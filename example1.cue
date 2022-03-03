package c4

elementsTags: aSupprimer: {legendText: "A Migrer en React", bgColor: "#00239c", borderColor: "#ff0239c", technology: React, shadowing: true, shape: "eightsided"}
relationsTags: myTest: {lineStyle: "dash"}

// Containers

myWebApp: #Container & {
	id:          "web_app"
	label:       "Web Application"
	description: "Allows users to compare multiple Twitter timelines"
	technology:  CSharp
	tags: [elementsTags.aSupprimer]
}

myothercontainer: #Container & {
	id:         "othercontainer"
	label:      "Cool"
	technology: React
}

twitter: #System & {
	id:    "twitter"
	label: "Twitter"
	link:  "https://www.twitter.com"
}

sampleSystem: #System & {
	id:    "c1"
	label: "Sample System"
	containers: [myWebApp, myothercontainer]
}

admin: #Person & {
	id:    "admin"
	label: "Administrator"
}

C1: #C1 & {
	Persons: [admin]
	Technologies: [ React, CSharp]
	Systems: [sampleSystem, twitter]
	Relations: [
		{source: admin, dest:    myWebApp, description: "Uses", protocol:            "HTTPS ", tags: [relationsTags.myTest]},
		{source: myWebApp, dest: twitter, description:  "Get tweets from", protocol: "HTTPS ", link: "https://plantuml.com/link"},
	]
}
