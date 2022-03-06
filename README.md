# cue4puml4c4

Heavy work in progress...

## Pure cue

- clone the repo
- in `example` dir, tweak `example1.cue`
- `cue vet` to check... and `cue genpuml` to generate the plantuml

## live view

- Be sure to have a plantuml server running at `localhost:8080`.
- within the `example` dir
- Run the local server: `go run ../cmd`
- open a browser to [http://localhost:9090](http://localhost:9090)
- save a file in the current directory
- profit :D

## Simple HOWTO

- Create a directory to create a new diagram
- create a cue file with your diagram
- import the `c4` package:

```cue
import "github.com/owulveryck/cue4puml4c4:c4"
```

- create a `C1` diagram of type `c4.#C1` and add at least one system

```cue
package main

import "github.com/owulveryck/cue4puml4c4:c4"

C1: c4.#C1 & {
 Systems: [{id: "sample", label: "my sample"}]
}

```

- create a tool to dump your `C1` object into plantuml (eg: `plantuml_tool.cue`):

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
