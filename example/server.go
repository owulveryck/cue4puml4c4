package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
)

type diagram struct {
	image    []byte
	plantuml []byte
}

func main() {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("NewWatcher failed: ", err)
	}

	err = watcher.Add("./")
	if err != nil {
		log.Fatal("Add failed:", err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%v", index)
	})
	http.HandleFunc("/connws/", (&client{
		upgrader: upgrader,
		watcher:  watcher,
		watch:    watch,
	}).ConnWs)
	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

type client struct {
	upgrader websocket.Upgrader
	watcher  *fsnotify.Watcher
	watch    func(ctx context.Context, watcher *fsnotify.Watcher, C chan diagram)
}

func (c *client) ConnWs(w http.ResponseWriter, r *http.Request) {
	C := make(chan diagram, 5)
	log.Printf("New Connection with %v", C)
	defer func() {
		// drain and close
		for len(C) > 0 {
			<-C
		}
		close(C)
	}()
	go c.watch(r.Context(), c.watcher, C)
	ws, err := c.upgrader.Upgrade(w, r, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		log.Println(err)
		return
	}

	res := map[string]interface{}{}
	for {
		select {
		case diagram := <-C:
			str := base64.StdEncoding.EncodeToString(diagram.image)
			res["image"] = str

			if err = ws.WriteJSON(&res); err != nil {
				log.Println(err)
			}
			log.Println("new picture sent")
		case <-r.Context().Done():
			return
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
