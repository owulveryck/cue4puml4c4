---
title: "#GetAllElementTags"
linkTitle: "#GetAllElementTags"
type: docs
description: >-
     This is a description of the #GetAllElementTags object
---

## Definition

```cue

_#def
_#def: {
	#next: _
	#func: {
		#in: _
		tags: {
			for i, x in #in {
				if x.tags != _|_ // explicit error (_|_ literal) in source
				{
					for tag in x.tags {
						"\(tag.id)": tag
					}
				}
				if x.systems != _|_ // explicit error (_|_ literal) in source
				{
					(#next & {
						#in: x.systems
					}).tags
				}
				if x.containers != _|_ // explicit error (_|_ literal) in source
				{
					(#next & {
						#in: x.containers
					}).tags
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
		tags: {
			for i, x in [#in](../in) {
				if x.tags != _|_ // explicit error (_|_ literal) in source
				{
					for tag in x.tags {
						"\(tag.id)": tag
					}
				}
				if x.systems != _|_ // explicit error (_|_ literal) in source
				{
					([#next](../next) & {
						[#in](../in): x.systems
					}).tags
				}
				if x.containers != _|_ // explicit error (_|_ literal) in source
				{
					([#next](../next) & {
						[#in](../in): x.containers
					}).tags
				}
			}
		}
	}

