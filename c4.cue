package c4

#System: {
	id:           #id
	label:        *id | string
	description?: string
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
