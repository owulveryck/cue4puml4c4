---
title: "#RelationTag"
linkTitle: "#RelationTag"
type: docs
description: >-
     This is a description of the #RelationTag object
---

## Definition

```cue

_#def
_#def: {
	id:            string
	textColor?:    #color
	lineColor?:    #color
	lineStyle?:    "dash" | "dot" | "bold"
	technology?:   #Technology
	legendText?:   string
	legendSprite?: "\(technology)"
}
```

## Usage
- id:            string
- textColor?:    [#color](../color)
- lineColor?:    [#color](../color)
- lineStyle?:    "dash" | "dot" | "bold"
- technology?:   [#Technology](../technology)
- legendText?:   string
- legendSprite?: "\(technology)"

