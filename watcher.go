package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func watch(ctx context.Context, watcher *fsnotify.Watcher, C chan diagram) {
	err := generateImageAndSend(C)
	if err != nil {
		log.Println(err)
	}
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op == fsnotify.Write {
				err := generateImageAndSend(C)
				if err != nil {
					log.Println(err)
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		case <-ctx.Done():
			return
		}
	}
}

func generateImageAndSend(C chan diagram) error {
	cue := exec.Command("cue", "cmd", "genpuml")
	var outb, errb bytes.Buffer
	cue.Stdout = &outb
	cue.Stderr = &errb
	if err := cue.Run(); err != nil {
		return fmt.Errorf("%v: %v", err, errb.String())
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
		return err
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
	puml, _ := format(bytes.NewReader(outb.Bytes()))
	log.Printf("%s", puml)
	C <- diagram{
		plantuml: outb.Bytes(),
		image:    res,
	}
	return nil
}

func format(r io.Reader) (string, error) {
	var b strings.Builder
	scanner := bufio.NewScanner(r)
	inBlock := 0
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}
		if text[0:1] == `/` {
			b.WriteRune('\n')
		}
		if text[len(text)-1:] == `}` {
			inBlock--
		}
		for i := 0; i < inBlock; i++ {
			b.WriteRune('\t')
		}
		b.WriteString(text)
		b.WriteRune('\n')
		if text[len(text)-1:] == `{` {
			inBlock++
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return b.String(), nil
}
