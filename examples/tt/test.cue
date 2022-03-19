package main

import (
	"github.com/owulveryck/cue4puml4c4:c4"
	"github.com/owulveryck/cue4puml4c4/tags/teamtopologies:tt"
	"github.com/owulveryck/cue4puml4c4/technology/dev"
	"github.com/owulveryck/cue4puml4c4/technology/stdlib"
)

t1: c4.#System & {
	id:    "t1"
	label: "Team1 / Devs"
	tags: [tt.streamAlignedTeam]
	containers: [{id: "code", label: "code", technology: dev.Java}]
	systems: [t2]
}
t2: c4.#System & {
	id:    "t2"
	label: "Team2 / Data Scientists"
	tags: [tt.complicatedSubsystemTeam]
}
t3: c4.#System & {
	id:    "t3"
	label: "Team3 / YAML Engineers"
	tags: [tt.enablingTeam]
	containers: [{id: "yaml", label: "configs", technology: stdlib.Yaml}]
}
t4: c4.#System & {
	id:    "t4"
	label: "Team4 / Ops"
	tags: [tt.platformTeam]
	containers: [{id: "k8s", label: "Kubernetes", technology: stdlib.Kubernetes}]
}
C1: c4.#C1 & {
	Systems: [t1, t3, t4]
}
