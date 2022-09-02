---
title: "#GetAllRelationTags"
linkTitle: "#GetAllRelationTags"
type: docs
description: >-
     This is a description of the #GetAllRelationTags object
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
				if (x & #Relation) != _|_ // explicit error (_|_ literal) in source
				{
					if x.tags != _|_ // explicit error (_|_ literal) in source
					{
						for t in x.tags {
							"\(t.id)": t
						}
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
				if x.relations != _|_ // explicit error (_|_ literal) in source
				{
					(#next & {
						#in: x.relations
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
				if (x & [#Relation](../relation)) != _|_ // explicit error (_|_ literal) in source
				{
					if x.tags != _|_ // explicit error (_|_ literal) in source
					{
						for t in x.tags {
							"\(t.id)": t
						}
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
				if x.relations != _|_ // explicit error (_|_ literal) in source
				{
					([#next](../next) & {
						[#in](../in): x.relations
					}).tags
				}
			}
		}
	}

