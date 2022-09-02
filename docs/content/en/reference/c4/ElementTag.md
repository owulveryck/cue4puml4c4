---
title: "#ElementTag"
linkTitle: "#ElementTag"
type: docs
description: >-
     This is a description of the #ElementTag object
---

## Definition

```cue

_#def
_#def: {
	id:           string
	bgColor?:     #color
	fontColor?:   #color
	borderColor?: #color
	legendText?:  string
	technology?:  #Technology
	shadowing?:   bool
	shape?:       "rounded" | "eightsided"
}
```

## Usage
- id:           string
- bgColor?:     [#color](../color)
- fontColor?:   [#color](../color)
- borderColor?: [#color](../color)
- legendText?:  string
- technology?:  [#Technology](../technology)
- shadowing?:   bool
- shape?:       "rounded" | "eightsided"

