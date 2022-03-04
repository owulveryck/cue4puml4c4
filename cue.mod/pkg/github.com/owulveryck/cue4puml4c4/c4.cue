package c4

#id: =~"^[a-zA-Z][a-zA-Z_]*"

#System: {
	label:        string
	id:           #id
	description?: string
	isBoundary:   bool | *false
	technology?:  #Technology
	type?:        #Type
	containers?: [...#Container]
	rels?: [...#Rel]
	systems?: [...#System]
	link?: #url
	tags?: [...#ElementTag]
}

#Container: {
	id:           #id
	label:        string
	technology:   #Technology | *noTech
	description?: string
	tags?: [...#ElementTag]
	link?: #url
}

#Person: {
	id:    #id
	label: string
}

#url: string

noTech: #Technology & {
	name: "Undefined"
}

#Rel: {
	source:      #System | #Container | #Person
	dest:        #System | #Container | #Person
	description: string | *""
	protocol?:   string
	link?:       #url
	tags?: [...#RelationTag]
}

#color: =~"#[0-9a-f]{6}"

#ElementTag: {
	id:           string
	bgColor?:     #color
	fontColor?:   #color
	borderColor?: #color
	legendText?:  string
	technology?:  #Technology
	shadowing?:   bool
	shape?:       "rounded" | "eightsided"
}

#RelationTag: {
	id:            string
	textColor?:    #color
	lineColor?:    #color
	lineStyle?:    "dash" | "dot" | "bold"
	technology?:   #Technology
	legendText?:   string
	legendSprite?: "\(technology)"
}

#Technology: {
	name: string
	type: *"" | "Db" | "Queue"
	sprite?: {
		id:  string | *"\(name)"
		url: string
	}
}

#Sprite: {
	name: string
	url:  #url
}

#Type: {
	name: string
}
#C1: {
	elementTags:  elementsTags
	relationTags: relationsTags
	Technologies?: [...#Technology]
	Persons?: [...#Person]
	Systems: [...#System]
	Container?: [...#Container]
	Relations?: [...#Rel]
}

// Tags
elementsTags: [ID=_]: #ElementTag & {
	id: "\(ID)"
}

relationsTags: [ID=_]: #RelationTag & {
	id: "\(ID)"
}
