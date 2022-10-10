---
title: Team Topologies tags
linkTitle: team-topologies
type: docs
menu:
  main:
    weight: 5

---

Here is an example with team topologies

{{< tabpane >}}
{{< tab header="Diagram data (CUE)" lang="cue" >}}
package main

import (
 "github.com/owulveryck/cue4puml4c4:c4"
 "github.com/owulveryck/cue4puml4c4/tags/teamtopologies:tt"
)

ptf: c4.#System & {
 id:    "ptf"
 label: "Platform team"
 description: """
  a grouping of other team types that provide a compelling internal product to accelerate delivery by Stream-aligned teams
  """
 link: "https://teamtopologies.com/key-concepts"
 tags: [tt.platformTeam]
}
streamAligned: c4.#System & {
 id:    "streamAligned"
 label: "Stream-aligned team"
 description: """
   aligned to a flow of work from (usually) a segment of the business domain
  """
 link: "https://teamtopologies.com/key-concepts"
 tags: [tt.streamAlignedTeam]
}

complicatedSubSystem: c4.#System & {
 id:    "complicatedSubSystem"
 label: "Complicated-subsystem team"
 description: """
   where significant mathematics/calculation/technical expertise is needed.
  """
 link: "https://teamtopologies.com/key-concepts"
 tags: [tt.complicatedSubsystemTeam]
}

enabling: c4.#System & {
 id:    "enabling"
 label: "Enabling Team"
 description: """
   helps a Stream-aligned team to overcome obstacles. Also detects missing capabilities.
  """
 link: "https://teamtopologies.com/key-concepts"
 tags: [tt.enablingTeam]
}

C1: c4.#C1 & {
 Systems: [ptf, complicatedSubSystem, streamAligned, enabling]
}
{{< /tab >}}
{{< tab header="CUE command" lang="cue" >}}
package main

import (
        "tool/cli"
        "text/template"
        "github.com/owulveryck/cue4puml4c4:c4"
)

command: genpuml: {
        c1: cli.Print & {
                text: template.Execute(c4.plantumlTemplate, C1) // change C1 here with the name of your object
        }
}
{{< /tab >}}
{{< tab header="Plantuml code" >}}
@startuml MyDiagram
@startuml MyDiagram
skinparam backgroundcolor transparent
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml
!include https://raw.githubusercontent.com/owulveryck/PlantUML-icons-GCP/master/official/GCPCommon.puml
LAYOUT_TOP_DOWN()
HIDE_STEREOTYPE()
/'Relation Tags'/ 
/'Element Tags'/ 
AddElementTag("complicatedSubsystemTeam",$bgColor="#f4c399",$borderColor="#d39450",$shape="EightSidedShape()",$legendText="Complicated-subsystem team")
AddElementTag("enablingTeam",$bgColor="#ad98c2",$borderColor="#6b5294",$shape="RoundedBoxShape()",$legendText="Enabling team")
AddElementTag("platformTeam",$bgColor="#9bcbe5",$borderColor="#6b9bc1",$shape="RoundedBoxShape()",$legendText="Platform team")
AddElementTag("streamAlignedTeam",$bgColor="#fae1a4",$borderColor="#f2c24f",$shape="RoundedBoxShape()",$legendText="Stream-aligned team")
	
/' Systems '/
System(ptf,"Platform team","a grouping of other team types that provide a compelling internal product to accelerate delivery by Stream-aligned teams",$link="https://teamtopologies.com/key-concepts",$tags="platformTeam+ ")
System(complicatedSubSystem,"Complicated-subsystem team"," where significant mathematics/calculation/technical expertise is needed.",$link="https://teamtopologies.com/key-concepts",$tags="complicatedSubsystemTeam+ ")
System(streamAligned,"Stream-aligned team"," aligned to a flow of work from (usually) a segment of the business domain",$link="https://teamtopologies.com/key-concepts",$tags="streamAlignedTeam+ ")
System(enabling,"Enabling Team"," helps a Stream-aligned team to overcome obstacles. Also detects missing capabilities.",$link="https://teamtopologies.com/key-concepts",$tags="enablingTeam+ ")
SHOW_LEGEND()
@enduml
{{< /tab >}}
{{< tab header="Plantuml rendering" lang="plantuml" >}}
@startuml MyDiagram
skinparam backgroundcolor transparent
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml
!include https://raw.githubusercontent.com/owulveryck/PlantUML-icons-GCP/master/official/GCPCommon.puml
LAYOUT_TOP_DOWN()
HIDE_STEREOTYPE()
/'Relation Tags'/ 
/'Element Tags'/ 
AddElementTag("complicatedSubsystemTeam",$bgColor="#f4c399",$borderColor="#d39450",$shape="EightSidedShape()",$legendText="Complicated-subsystem team")
AddElementTag("enablingTeam",$bgColor="#ad98c2",$borderColor="#6b5294",$shape="RoundedBoxShape()",$legendText="Enabling team")
AddElementTag("platformTeam",$bgColor="#9bcbe5",$borderColor="#6b9bc1",$shape="RoundedBoxShape()",$legendText="Platform team")
AddElementTag("streamAlignedTeam",$bgColor="#fae1a4",$borderColor="#f2c24f",$shape="RoundedBoxShape()",$legendText="Stream-aligned team")
	
/' Systems '/
System(ptf,"Platform team","a grouping of other team types that provide a compelling internal product to accelerate delivery by Stream-aligned teams",$link="https://teamtopologies.com/key-concepts",$tags="platformTeam+ ")
System(complicatedSubSystem,"Complicated-subsystem team"," where significant mathematics/calculation/technical expertise is needed.",$link="https://teamtopologies.com/key-concepts",$tags="complicatedSubsystemTeam+ ")
System(streamAligned,"Stream-aligned team"," aligned to a flow of work from (usually) a segment of the business domain",$link="https://teamtopologies.com/key-concepts",$tags="streamAlignedTeam+ ")
System(enabling,"Enabling Team"," helps a Stream-aligned team to overcome obstacles. Also detects missing capabilities.",$link="https://teamtopologies.com/key-concepts",$tags="enablingTeam+ ")
SHOW_LEGEND()
@enduml
{{< /tab >}}

{{< /tabpane >}}

