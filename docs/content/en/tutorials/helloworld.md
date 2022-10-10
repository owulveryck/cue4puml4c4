---
title: "Hello World"
linkTitle: "Hello World"
type: docs
weight: 10
description: >-
     Your first diagram from cue to plantuml
---

In this tutorial you will:

- represent a very simple C1 diagram as cue data
- render it into plantuml

This tutorial works on MacOS and Linux, but should be adaptable to Windows easily.

## Prerequisites

Only the `cue` utility is required. You can download is from [this page](https://github.com/cue-lang/cue/releases/)

## Setting up the directory

```shell
mkdir $TMPDIR/helloworld
cd $TMPDIR/helloworld
```

## Bootstraping the project

Init a new cue module:

```shell
cue mod init
```

The vendor the cue4puml4c4 library:

```shell
mkdir -p cue.mod/pkg/github.com/owulveryck/cue4puml4c4
git clone https://github.com/owulveryck/cue4puml4c4.git cue.mod/pkg/github.com/owulveryck/cue4puml4c4
```

## Create the command to generate the puml locally

Write the content below in a file name `command_tool.cue` with your favorite editor:

```cue
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
```

## Create the diagram

in a file named `test.cue` write the content below:

```cue
package main

import "github.com/owulveryck/cue4puml4c4:c4"

C1: c4.#C1 & { // the name C1 should be coherent with the name you declare in the command
        Systems: [{id: "sample", label: "my sample"}]
}
```

## Generate the plantuml diagram

run the following command to generate the plantuml diagram:

```shell
cue genpuml | grep -v '^$'
```

_Note_: the `| grep` command is to remove the blank lines for brevity but does not impact the result.

Generates the following output:

```puml
@startuml MyDiagram
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml
!include https://raw.githubusercontent.com/owulveryck/PlantUML-icons-GCP/master/official/GCPCommon.puml
LAYOUT_TOP_DOWN()
HIDE_STEREOTYPE()
/'Relation Tags'/ 
/'Element Tags'/ 
/' Systems '/
System(sample,"my sample")
SHOW_LEGEND()
@enduml
```

Which renders as:
```plantuml
@startuml MyDiagram
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml
!include https://raw.githubusercontent.com/owulveryck/PlantUML-icons-GCP/master/official/GCPCommon.puml
LAYOUT_TOP_DOWN()
HIDE_STEREOTYPE()
/'Relation Tags'/ 
/'Element Tags'/ 
/' Systems '/
System(sample,"my sample")
SHOW_LEGEND()
@enduml
```

You can copy/paste the content in [plantuml server online](https://www.plantuml.com/plantuml/uml/bSt1Qy8m5CVnU_-Ap4iLrkQoqqwEDkYWiyLMn5E9QTImvAKaBvF-zuiKOTVT7Xzul-_Q8tSOZ4vU50WT1abkYAD_fzJnUtvLU0lduAKJ5b02vi8QEceEZv_Cuw3LcHPlMVEWmrT53VSeNQI-i81SWNJv-CzzUm_wHxf9VDFdRgQ41PzjYkPPi9UB4efh6gsmnbYOEzN7gJvqVLStVLaVTyjLilqiMTzsRC_gxjImA3JTIyrHMI0Tl_gKaaXCIoC1_mXffvWodgGqcS-bvsRKydLX9ZAVYrNIRkjZNx4DsvLnUorX24R_0W00) to see how it renders

## Cleaning

```shell
rm $TMPDIR/helloworld/*.cue
rmdir $TMPDIR/helloworld
```
