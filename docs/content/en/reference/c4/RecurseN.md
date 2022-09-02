---
title: "#RecurseN"
linkTitle: "#RecurseN"
type: docs
description: >-
     This is a description of the #RecurseN object
---

## Definition

```cue
import "list"

_#def
_#def: {
	#maxiter: int & >=0 | *8
	#funcFactory: {
		#next: _
		#func: _
	}
	for k, v in list.Range(0, #maxiter, 1) {
		// this is where we build up our indexed functions and the references between them
		#funcs: {
			"\(k)": (#funcFactory & {
				#next: #funcs["\(k+1)"]
			}).#func
		}
	}
	#funcs: {
		"\(#maxiter)": null
	}
	#funcs["0"]
}
```

## Usage
- [#maxiter](../maxiter): int & >=0 | *8
- [#funcFactory](../funcfactory): {
		[#next](../next): _
		[#func](../func): _
	}
- for k, v in list.Range(0, [#maxiter](../maxiter), 1) {
		// this is where we build up our indexed functions and the references between them
		[#funcs](../funcs): {
			"\(k)": ([#funcFactory](../funcfactory) & {
				[#next](../next): #funcs["\(k+1)"]
			}).[#func](../func)
		}
	}
- [#funcs](../funcs): {
		"\([#maxiter](../maxiter))": null
	}
- [#funcs](../funcs)["0"]

