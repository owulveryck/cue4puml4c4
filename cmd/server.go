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
			res["img"] = diagram.image
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
	index = `<!DOCTYPE html>
	<html lang="en-EN">
	
	<head>
	  <link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.4.0/css/font-awesome.min.css" rel="stylesheet" type="text/css" />
	  <meta charset="utf-8">
	  <style>
	    html,
	    body {
	      height: 100%;
	      margin: 0;
	      background-color: gray;
	    }
	    /* Make the download icon big. */
	    .download-button{
	      font-size: 32em;
	      text-align: center;
	    }
      
	    /* Make the download icon look clickable when you hover over it. */
	    .download-button i {
	      cursor: pointer;
	    }
	
	    .container {
	      max-width: 100%;
	      max-height: 100%;
	      bottom: 0;
	      left: 0;
	      margin: auto;
	      overflow: auto;
	      position: fixed;
	      right: 0;
	      top: 0;
	      -o-object-fit: contain;
	      object-fit: contain;
	    }
	
	    .right {
	      float: right;
	      width: 80%;
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
	
	    object {
	      aspect-ratio: inherit;
	      max-height: 100%;
	      max-width: 100%;
	
	    }
	  </style>
	  <title>{{.Path}}</title>
	</head>
	
	<body>
	  <div class="container">
	    <object id="output" type="image/svg+xml" data="">Content</object>
	</div>
	</body>
	
	<script type="text/javascript">
	  var url = "ws{{if eq .Scheme "https"}}s{{end}}://{{.Host}}{{.Path}}";
	  ws = new WebSocket(url);
	
	  ws.onopen = function () {
	    console.log("[onopen] connect ws uri.");
	    var data = {
	      "Action": "requireConnect"
	    };
	    ws.send(JSON.stringify(data));
	  }
	
	  ws.onmessage = function (e) {
	    console.log("[onmessage] receive message.");
	    var res = JSON.parse(e.data);
	    document.getElementById("output").setAttribute("data", "data:image/svg+xml;utf8;base64," + res["image"]); // decodeURIComponent(escape(window.atob(res["image"]))))
	    document.getElementById("dl").setAttribute("href", "data:image/svg+xml;utf8;base64," + res["image"]); // decodeURIComponent(escape(window.atob(res["image"]))))
	    console.log(res);
	  }
	
	  ws.onclose = function (e) {
	    console.log("[onclose] connection closed (" + e.code + ")");
	  }
	
	  ws.onerror = function (e) {
	    console.log("[onerror] error!");
	  }
	</script>
	
	</html>
	`
)
