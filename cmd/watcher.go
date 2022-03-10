package main

import (
	"context"
	"log"
	"regexp"

	"github.com/fsnotify/fsnotify"
)

// watch for the event on watcher, then calls the generateImageAndSend
func watch(ctx context.Context, watcher watchedDir, C chan diagram, errorC chan error) {
	err := generateImageAndSend(watcher.path, C)
	if err != nil {
		errorC <- err
	}
	re := regexp.MustCompile(`\.cue$`)
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			isCue := re.MatchString(event.Name)
			if (event.Op == fsnotify.Write || event.Op == fsnotify.Create) && isCue {
				log.Printf("%v modified", event.Name)
				err := generateImageAndSend(watcher.path, C)
				if err != nil {
					errorC <- err
				}
			}
		case err, ok := <-watcher.Errors:
			errorC <- err
			if !ok {
				return
			}
		case <-ctx.Done():
			return
		}
	}
}
