---
title: "#Technology"
linkTitle: "#Technology"
type: docs
description: >-
     This is a description of the #Technology object
---

## Definition

```cue

_#def
_#def: {
	name: string
	type: *"" | "Db" | "Queue"
	sprite?: {
		id:  string | *"\(name)"
		url: string
	}
}
```

## Usage
- name: string
- type: *"" | "Db" | "Queue"
- sprite?: {
		id:  string | *"\(name)"
		url: string
	}

