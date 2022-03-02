package c4

#System: {
	label:        string
	id:           string
	description?: string
	isBoundary:   bool | *false
	technology?:  #Technology
	type?:        #Type
	containers?: [...#Container]
	rels?: [...#Rel]
	systems?: [...#System]
}

#Container: {
	id:           string
	label:        string
	technology:   #Technology | *noTech
	description?: string
	tags?: [...#tag]
	link?: #url
}

#tag: string

#url: string

noTech: #Technology & {
	name: "Undefined"
}

#Rel: {
	source:      #System | #Container
	dest:        #System | #Container
	description: string | *""
	protocol?:   string
	link?:       #url
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
	Technologies?: [...#Technology]
	Systems: [...#System]
	Container?: [...#Container]
}
