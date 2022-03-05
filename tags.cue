package c4

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

// Tags
elementsTags: [ID=_]: #ElementTag & {
	id: "\(ID)"
}

relationsTags: [ID=_]: #RelationTag & {
	id: "\(ID)"
}
