package c4

import (
	"list"

)

#RecurseN: {
	// this is the bound on our recursion
	#maxiter: uint | *8

	// This is the function list element
	// we generate this to simulate recursion
	#funcFactory: {
		#next: _
		#func: _
	}

	// this is our "recursion unrolling"
	for k, v in list.Range(0, #maxiter, 1) {
		// this is where we build up our indexed functions and the references between them
		#funcs: "\(k)": (#funcFactory & {#next: #funcs["\(k+1)"]}).#func
	}

	// our final value needs to be null
	#funcs: "\(#maxiter)": null

	// we embed the head of the list so that the values
	// we write this into can be used like other CUE "functions"
	#funcs["0"]
}
#GetAllTech: {
	#next: _
	#func: {
		#in: _
		tech: {
			for i, x in #in {
				if x.technology != _|_ {
					"\(x.technology.name)": x.technology
				}
				if x.systems != _|_ {
					(#next & {#in: x.systems}).tech
				}
				if x.containers != _|_ {
					(#next & {#in: x.containers}).tech
				}
			}
		}
	}
}

#GetAllElementTags: {
	#next: _
	#func: {
		#in: _
		tags: {
			for i, x in #in {
				if x.tags != _|_ {
					for tag in x.tags {
						"\(tag.id)": tag
					}
				}
				if x.systems != _|_ {
					(#next & {#in: x.systems}).tags
				}
				if x.containers != _|_ {
					(#next & {#in: x.containers}).tags
				}
			}
		}
	}
}

#GetAllRelationTags: {
	#next: _
	#func: {
		#in: _
		tags: {
			for i, x in #in {
				if (x & #Relation) != _|_ {
					if x.tags != _|_ {
						for t in x.tags {
							"\(t.id)": t
						}
					}
				}
				if x.systems != _|_ {
					(#next & {#in: x.systems}).tags
				}
				if x.containers != _|_ {
					(#next & {#in: x.containers}).tags
				}
				if x.relations != _|_ {
					(#next & {#in: x.relations}).tags
				}
			}
		}
	}
}

#FoundTechs:        #RecurseN & {#funcFactory: #GetAllTech}
#FoundElementTags:  #RecurseN & {#funcFactory: #GetAllElementTags}
#FoundRelationTags: #RecurseN & {#funcFactory: #GetAllRelationTags}
