---
title: "#Container"
linkTitle: "#Container"
type: docs
description: >-
     This is a description of the #Container object
---

## Definition

```cue
import "strings"

_#def
_#def: {
	id:           #id
	label:        *id | string
	technology:   #Technology | *_noTech
	description?: string
	desc:         "" | *strings.Replace(description, """


		""", "\\n", -1)
	tags?: [...#ElementTag]
	link?: #url
}
```

## Usage
- id:           [#id](../id)
- label:        *id | string
- technology:   [#Technology](../technology) | *_noTech
- description?: string
- desc:         "" | *strings.Replace(description, """
		""", "\\n", -1)
- tags?: [...[#ElementTag](../elementtag)]
- link?: [#url](../url)

