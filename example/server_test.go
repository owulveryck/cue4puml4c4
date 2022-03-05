package main

import (
	"bytes"
	"testing"
)

func TestFormat(t *testing.T) {
	out, err := format(bytes.NewBuffer([]byte(sample)))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)
}

const (
	expected = `@startuml MyName
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml
!include https://raw.githubusercontent.com/owulveryck/PlantUML-icons-GCP/master/official/GCPCommon.puml
!include https://raw.githubusercontent.com/tupadr3/plantuml-icon-font-sprites/master/devicons/react.puml 
!include https://raw.githubusercontent.com/tupadr3/plantuml-icon-font-sprites/master/devicons2/csharp.puml 
AddRelTag("myTest",$lineStyle=DashedLine())

/'Element Tags'/ 
AddElementTag("aSupprimer",$bgColor="#00239c",$borderColor="#ff0239c",$shadowing="true",$shape="EightSidedShape()",$legendText="A Migrer en React",$techn="Typescript/React",$sprite="react")
Person(admin,"Administrator")

System(c1,"Sample System"){
	Container(web_app,"Web Application","C#","Allows users to compare multiple Twitter timelines","csharp",$tags="aSupprimer+ ")
	Container(othercontainer,"Cool","Typescript/React","","react")
}

System(twitter,"Twitter",$link="https://www.twitter.com")
Rel("admin","web_app","Uses","HTTPS ",$tags="myTest+ ")
Rel("web_app","twitter","Get tweets from","HTTPS ",$link="https://plantuml.com/link")
SHOW_LEGEND()
@enduml
`
	sample = `
@startuml MyName
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml


!include https://raw.githubusercontent.com/owulveryck/PlantUML-icons-GCP/master/official/GCPCommon.puml


!include https://raw.githubusercontent.com/tupadr3/plantuml-icon-font-sprites/master/devicons/react.puml 
!include https://raw.githubusercontent.com/tupadr3/plantuml-icon-font-sprites/master/devicons2/csharp.puml 
AddRelTag("myTest",$lineStyle=DashedLine())

/'Element Tags'/ 



AddElementTag("aSupprimer",$bgColor="#00239c",$borderColor="#ff0239c",$shadowing="true",$shape="EightSidedShape()",$legendText="A Migrer en React",$techn="Typescript/React",$sprite="react")
Person(admin,"Administrator")
System(c1,"Sample System"){
	Container(web_app,"Web Application","C#","Allows users to compare multiple Twitter timelines","csharp",$tags="aSupprimer+ ")
	Container(othercontainer,"Cool","Typescript/React","","react")
}

System(twitter,"Twitter",$link="https://www.twitter.com")


Rel("admin","web_app","Uses","HTTPS ",$tags="myTest+ ")
Rel("web_app","twitter","Get tweets from","HTTPS ",$link="https://plantuml.com/link")
SHOW_LEGEND()
@enduml
`
)
