---
title: "#System"
linkTitle: "#System"
type: docs
description: >-
     This is a description of the #System object
---

## Definition

```cue
import "strings"

_#def
_#def: {
	id:           #id
	label:        *id | string
	description?: string
	desc:         "" | *strings.Replace(description, """


		""", "\\n", -1)
	isBoundary:   bool | *false
	technology?:  #Technology
	type?:        #Type
	containers?: [...#Container]
	relations?: [...#Relation]
	systems?: [...#System]
	link?: #url
	tags?: [...#ElementTag]
}
```

## Usage
- id:           [#id](../id)
- label:        *id | string
- description?: string
- desc:         "" | *strings.Replace(description, """
		""", "\\n", -1)
- isBoundary:   bool | *false
- technology?:  [#Technology](../technology)
- type?:        [#Type](../type)
- containers?: [...[#Container](../container)]
- relations?: [...[#Relation](../relation)]
- systems?: [...[#System](../system)]
- link?: [#url](../url)
- tags?: [...[#ElementTag](../elementtag)]

