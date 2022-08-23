# Team Topologies tags

usage:

```cue
package main

import (
	"github.com/owulveryck/cue4puml4c4:c4"
	"github.com/owulveryck/cue4puml4c4/tags/teamtopologies:tt"
)

ptf: c4.#System & {
	id:          "ptf"
	label:       "ptf"
	description: "ptf"
	tags: [tt.platformTeam]
}
streamAligned: c4.#System & {
	id:          "streamAligned"
	label:       "streamAligned"
	description: "streamAligned"
	tags: [tt.streamAlignedTeam]
}

complicatedSubSystem: c4.#System & {
	id:          "complicatedSubSystem"
	label:       "complicatedSubSystem"
	description: "complicatedSubSystem"
	tags: [tt.complicatedSubsystemTeam]
}

enabling: c4.#System & {
	id:          "enabling"
	label:       "enabling"
	description: "enabling"
	tags: [tt.enablingTeam]
}

C1: c4.#C1 & {
	Systems: [ptf, complicatedSubSystem, streamAligned, enabling]
}
```

![Result](tt.svg)

