package main

import (
	"html/template"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
)

func processConfig(config configuration) (*http.ServeMux, error) {
	mux := http.NewServeMux()
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	err = watcher.Add(config.PollingDir)
	if err != nil {
		return nil, err
	}
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	tmpl, err := template.New("index").Parse(index)
	if err != nil {
		return nil, err
	}
	currentContext := &ctx{
		mux:           mux,
		upgrader:      upgrader,
		indexTemplate: tmpl,
		cwd:           cwd,
	}
	err = filepath.WalkDir(config.PollingDir, currentContext.process)
	return mux, err
}

type ctx struct {
	mux           *http.ServeMux
	upgrader      websocket.Upgrader
	indexTemplate *template.Template
	cwd           string
}

func (ctx *ctx) process(p string, d fs.DirEntry, err error) error {
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
		return err
	}

	err = watcher.Add(filepath.Clean(filepath.Join(ctx.cwd, p)))
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
	files, err := ioutil.ReadDir(p)
	if err != nil {
		return err
	}
	type dirEntry struct {
		Path string
		Name string
	}
	dirEntries := make([]dirEntry, 0)
	for _, f := range files {
		if f.IsDir() {
			dirEntries = append(dirEntries, dirEntry{
				Name: f.Name(),
				Path: path.Clean(path.Join(cleanP, f.Name())),
			})
		}
	}
	ctx.mux.HandleFunc(cleanP, func(w http.ResponseWriter, r *http.Request) {
		err := ctx.indexTemplate.Execute(w, struct {
			URL        *url.URL
			Name       string
			DirEntries []dirEntry
		}{
			DirEntries: dirEntries,
			Name:       d.Name(),
			URL:        me,
		})
		if err != nil {
			log.Fatal(err)
		}

	})
	ctx.mux.HandleFunc(path.Join(cleanP, "connws/"), (&connexionHandler{
		upgrader: ctx.upgrader,
		watcher: watchedDir{
			Watcher: watcher,
			path:    p,
		},
		watch: watch,
	}).ConnWs)
	return nil
}
