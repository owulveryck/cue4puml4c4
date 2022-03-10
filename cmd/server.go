package main

import (
	"context"
	"encoding/base64"
	"log"
	"net/http"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
)

type diagram struct {
	svgImage []byte
	cue      []byte
	cueErr   []byte
	plantuml []byte
}

type connexionHandler struct {
	upgrader websocket.Upgrader
	watcher  watchedDir
	watch    func(ctx context.Context, watcher watchedDir, C chan diagram, errC chan error)
}

type watchedDir struct {
	path string
	*fsnotify.Watcher
}

func (c *connexionHandler) ConnWs(w http.ResponseWriter, r *http.Request) {
	C := make(chan diagram, 5)
	errC := make(chan error, 5)
	log.Printf("New Connection with %v", C)
	defer func() {
		// drain and close
		for len(C) > 0 {
			<-C
		}
		close(C)
	}()
	go c.watch(r.Context(), c.watcher, C, errC)
	ws, err := c.upgrader.Upgrade(w, r, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	res := map[string]interface{}{}
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		select {
		case diagram := <-C:
			str := base64.StdEncoding.EncodeToString(diagram.svgImage)
			res["image"] = str
			res["cue"] = diagram.cueErr
			res["plantuml"] = diagram.plantuml

			if err = ws.WriteJSON(&res); err != nil {
				log.Println(err)
			}
			log.Println("new picture sent")
		case err := <-errC:
			log.Println("error receiver:", err)
			res["error"] = []byte(err.Error())
			if err = ws.WriteJSON(&res); err != nil {
				log.Println(err)
			}
			delete(res, "error")
		case <-r.Context().Done():
			return
		}
	}
}
