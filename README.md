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

### Create your own diagrams

```shell
$ cue mod init
$ mkdir -p cue.mod/pkg/github.com/owulveryck/cue4puml4c4
$ git clone https://github.com/owulveryck/cue4puml4c4.git cue.mod/pkg/github.com/owulveryck/cue4puml4c4
$ cat <<EOF > test.cue
package main

import "github.com/owulveryck/cue4puml4c4:c4"

C1: c4.#C1 & { // the name C1 should be coherent with the name you declare in the command
        Systems: [{id: "sample", label: "my sample"}]
}
EOF

$ cat <<EOF > command_tool.cue
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
EOF
```
