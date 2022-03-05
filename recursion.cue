package main

import (
	"list"

)

#RecurseN: {
	// this is the bound on our recursion
	#maxiter: uint | *4

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
#Depth: #RecurseN & {#funcFactory: #GetAllTech}

techs: #Depth & {#in: C1.Systems}
