---
title: "#Relation"
linkTitle: "#Relation"
type: docs
description: >-
     This is a description of the #Relation object
---

## Definition

```cue
import "strings"

_#def
_#def: {
	source:      #System | #Container | #Person
	dest:        #System | #Container | #Person
	description: string | *""
	desc:        "" | *strings.Replace(description, """


		""", "\\n", -1)
	protocol?:   string
	link?:       #url
	tags?: [...#RelationTag]
}
```

## Usage
- source:      [#System](../system) | #Container | #Person
- dest:        [#System](../system) | #Container | #Person
- description: string | *""
- desc:        "" | *strings.Replace(description, """
		""", "\\n", -1)
- protocol?:   string
- link?:       [#url](../url)
- tags?: [...[#RelationTag](../relationtag)]

