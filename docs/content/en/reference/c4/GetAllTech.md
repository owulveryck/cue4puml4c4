---
title: "#GetAllTech"
linkTitle: "#GetAllTech"
type: docs
description: >-
     This is a description of the #GetAllTech object
---

## Definition

```cue

_#def
_#def: {
	#next: _
	#func: {
		#in: _
		tech: {
			for i, x in #in {
				if x.technology != _|_ // explicit error (_|_ literal) in source
				{
					"\(x.technology.name)": x.technology
				}
				if x.systems != _|_ // explicit error (_|_ literal) in source
				{
					(#next & {
						#in: x.systems
					}).tech
				}
				if x.containers != _|_ // explicit error (_|_ literal) in source
				{
					(#next & {
						#in: x.containers
					}).tech
				}
			}
		}
	}
}
```

## Usage
- [#next](../next): _
- [#func](../func): {
		[#in](../in): _
		tech: {
			for i, x in [#in](../in) {
				if x.technology != _|_ // explicit error (_|_ literal) in source
				{
					"\(x.technology.name)": x.technology
				}
				if x.systems != _|_ // explicit error (_|_ literal) in source
				{
					([#next](../next) & {
						[#in](../in): x.systems
					}).tech
				}
				if x.containers != _|_ // explicit error (_|_ literal) in source
				{
					([#next](../next) & {
						[#in](../in): x.containers
					}).tech
				}
			}
		}
	}

