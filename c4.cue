package c4

import "strings"

#System: {
	id:           #id
	label:        *id | string
	description?: string
	desc:         "" | *strings.Replace(description, "\n", "\\n", -1)
	isBoundary:   bool | *false
	technology?:  #Technology
	type?:        #Type
	containers?: [...#Container]
	relations?: [...#Relation]
	systems?: [...#System]
	link?: #url
	tags?: [...#ElementTag]
}

#Container: {
	id:           #id
	label:        *id | string
	technology:   #Technology | *_noTech
	description?: string
	desc:         "" | *strings.Replace(description, "\n", "\\n", -1)
	tags?: [...#ElementTag]
	link?: #url
}

#Person: {
	id:    #id
	label: *id | string
}

#Relation: {
	source:      #System | #Container | #Person
	dest:        #System | #Container | #Person
	description: string | *""
	desc:        "" | *strings.Replace(description, "\n", "\\n", -1)
	protocol?:   string
	link?:       #url
	tags?: [...#RelationTag]
}

#C1: {
	title?:      string
	elementTags: #FoundElementTags & {#in: Systems}
	relationTags: {a: {if relations != _|_ {#FoundRelationTags & {#in: relations}}}, b: #FoundRelationTags & {#in: Systems}}
	Technologies: #FoundTechs & {#in: Systems}
	Persons?: [...#Person]
	Systems: [...#System]
	SystemsExt?: [...#System]
	Container?: [...#Container]
	relations?: [...#Relation]
	layout:         *"top-down" | "left-right" | "landscape"
	hideStereotype: bool | *true
}
