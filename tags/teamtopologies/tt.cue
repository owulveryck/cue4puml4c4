package tt

// usage: import "github.com/owulveryck/cue4puml4c4/tags/teamtopologies:tt"

import (
	"github.com/owulveryck/cue4puml4c4:c4"
)

// Team Topologies
platformTeam: c4.#ElementTag & {
	id:          "platformTeam"
	legendText:  "Platform team"
	shape:       "rounded"
	bgColor:     "#9bcbe5"
	borderColor: "#6b9bc1"
}

streamAlignedTeam: c4.#ElementTag & {
	id:          "streamAlignedTeam"
	legendText:  "Stream-aligned team"
	shape:       "rounded"
	bgColor:     "#fae1a4"
	borderColor: "#f2c24f"
}

complicatedSubsystemTeam: c4.#ElementTag & {
	id:          "complicatedSubsystemTeam"
	legendText:  "Complicated-subsystem team"
	shape:       "eightsided"
	bgColor:     "#f4c399"
	borderColor: "#d39450"
}
enablingTeam: c4.#ElementTag & {
	id:          "enablingTeam"
	legendText:  "Enabling team"
	shape:       "rounded"
	bgColor:     "#ad98c2"
	borderColor: "#6b5294"
}
