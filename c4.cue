package c4

#System: {
	label:        string
	id:           #id
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
	label:        string
	technology:   #Technology | *_noTech
	description?: string
	tags?: [...#ElementTag]
	link?: #url
}

#Person: {
	id:    #id
	label: string
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
	elementTags:  #elementsTags
	relationTags: #relationsTags
	Technologies: #Depth & {#in: Systems}
	Persons?: [...#Person]
	Systems: [...#System]
	Container?: [...#Container]
	Relations?: [...#Relation]
}

// Tags
#elementsTags: [ID=_]: #ElementTag & {
	id: "\(ID)"
}

#relationsTags: [ID=_]: #RelationTag & {
	id: "\(ID)"
}
