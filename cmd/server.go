package main

import (
	"context"
	"encoding/base64"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"path"
	"path/filepath"
	"regexp"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
	"github.com/kelseyhightower/envconfig"
)

var config configuration

func init() {
	err := envconfig.Process("CUEWATCH", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
}

type configuration struct {
	PlantUMLServerAddress string `envconfig:"PLANTML_SERVER" default:"http://localhost:8080"`
	ListenAddress         string `envconfig:"ADDR" default:":9090"`
	PollingDir            string `envconfig:"POLLING_DIR" default:"./"`
	RecursivePoll         bool   `envconfig:"RECURSIVE" default:"true"`
}

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

	err = watcher.Add(config.PollingDir)
	if err != nil {
		log.Fatal("Add failed:", err)
	}
	if config.RecursivePoll {
		tmpl, err := template.New("index").Parse(index)
		if err != nil {
			log.Fatal(err)
		}
		err = filepath.WalkDir(config.PollingDir, func(p string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() {
				return nil
			}
			if isGit, err := regexp.MatchString(".*/.git/?.*", p); isGit || err != nil {
				return nil
			}
			watcher, err := fsnotify.NewWatcher()
			if err != nil {
				log.Fatal("NewWatcher failed: ", err)
			}

			err = watcher.Add(p)
			if err != nil {
				return err
			}
			cleanP := path.Clean(path.Join("/", p))
			log.Printf("Registering %v", cleanP)
			me, err := url.Parse("http://" + config.ListenAddress)
			if err != nil {
				return err
			}
			if me.Hostname() == "" {
				me.Host = "localhost" + me.Host
				me.Path = path.Join(cleanP, "/connws/")
			}
			http.HandleFunc(cleanP, func(w http.ResponseWriter, r *http.Request) {
				log.Println(me)
				tmpl.Execute(w, me)
			})
			http.HandleFunc(path.Join(cleanP, "connws/"), (&client{
				upgrader: upgrader,
				watcher: watchedDir{
					Watcher: watcher,
					path:    p,
				},
				watch: watch,
			}).ConnWs)
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		panic("todo")
	}
	err = http.ListenAndServe(config.ListenAddress, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

type client struct {
	upgrader websocket.Upgrader
	watcher  watchedDir
	watch    func(ctx context.Context, watcher watchedDir, C chan diagram)
}

type watchedDir struct {
	path string
	*fsnotify.Watcher
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
.svg-container {
	display: inline-block;
	width: 100%;
	position: relative;
	padding-bottom: 70%;
	vertical-align: middle;
	overflow: auto;
    }
    
</style>
<title>{{.Path}}</title>
</head>
<body>
<div class="container center">
<div id="output" class="svg-container"> </div>
</div>
</body>
</html>
<script type="text/javascript">
var url = "ws{{if eq .Scheme "https"}}s{{end}}://{{.Host}}{{.Path}}";
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
    document.getElementById("output").innerHTML = decodeURIComponent(escape(window.atob( res["image"] )))
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
