package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os/exec"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
)

var (
	upgrader websocket.Upgrader
	C        chan diagram
)

type diagram struct {
	image    []byte
	cue      []byte
	plantuml []byte
}

func main() {
	C = make(chan diagram)
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("NewWatcher failed: ", err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func(C chan diagram) {
		defer close(done)

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op == fsnotify.Write {
					cue := exec.Command("cue", "cmd", "genpuml")
					var outb, errb bytes.Buffer
					cue.Stdout = &outb
					cue.Stderr = &errb
					if err := cue.Run(); err != nil {
						log.Println(errb.String())
						log.Println(err)
						continue
					}
					u, _ := url.Parse("http://localhost:8080/plantuml/svg/")
					req := &http.Request{
						Method: http.MethodPost,
						URL:    u,
						Header: map[string][]string{
							"Content-Type": {"text/plain"},
						},
						Body: io.NopCloser(bytes.NewReader(outb.Bytes())),
					}
					resp, err := http.DefaultClient.Do(req)
					if err != nil {
						log.Println(err)
						log.Println("is plantuml server up and running?")
						continue
					}
					res, err := ioutil.ReadAll(resp.Body)

					if err != nil {
						log.Fatal(err)
					}
					resp.Body.Close()
					if resp.StatusCode != http.StatusOK {
						log.Println(resp.StatusCode)
					}
					log.Println("Senging new image")
					log.Printf("%s", outb.Bytes())
					C <- diagram{
						plantuml: outb.Bytes(),
						image:    res,
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}

	}(C)

	err = watcher.Add("./")
	if err != nil {
		log.Fatal("Add failed:", err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%v", index)
	})
	http.HandleFunc("/connws/", ConnWs)
	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func ConnWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		log.Println(err)
		return
	}

	res := map[string]interface{}{}
	for {
		if err = ws.ReadJSON(&res); err != nil {
			if err.Error() == "EOF" {
				return
			}
			// ErrShortWrite means a write accepted fewer bytes than requested then failed to return an explicit error.
			if err.Error() == "unexpected EOF" {
				return
			}
			fmt.Println("Read : " + err.Error())
			return
		}
		for diagram := range C {
			str := base64.StdEncoding.EncodeToString(diagram.image)
			res["image"] = str

			if err = ws.WriteJSON(&res); err != nil {
				log.Println(err)
			}
		}
	}
}

const (
	index = `
<html lang="zh-TW">
<head>
<style>
html, body {
  height: 100%;
  margin: 0;
  background-color: gray;
}
.container {
  margin: 0;
  display: block;
  height: 100%;
}
.center {
  margin: auto;
}
.container img {
  display: block;
  margin-left: auto;
  margin-right: auto;
  max-height: 100%;
  max-width: 100%;
}
</style>
</head>
<body>
<div class="container center">
<img id="image" />
</div>
</body>
</html>
<script type="text/javascript" src="http://code.jquery.com/jquery-1.11.1.min.js"></script>
<script type="text/javascript">
var url = "ws://localhost:9090/connws/";
ws = new WebSocket(url);

ws.onopen = function() {
  console.log("[onopen] connect ws uri.");
  var data = {
    "Action" : "requireConnect"
  };
  ws.send(JSON.stringify(data));
}

ws.onmessage = function(e) {
    console.log("[onmessage] receive message.");
    var res = JSON.parse(e.data);
    $("#image").attr("src", "data:image/svg+xml;utf8;base64," + res["image"]);
    console.log(res)
    console.log(res["cue"])
    console.log(res["plantuml"])
}

ws.onclose = function(e) {
    console.log("[onclose] connection closed (" + e.code + ")");
}

ws.onerror = function (e) {
    console.log("[onerror] error!");
}
</script>`
)
