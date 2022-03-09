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
	"path"
	"regexp"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func watch(ctx context.Context, watcher watchedDir, C chan diagram) {
	err := generateImageAndSend(watcher.path, C)
	if err != nil {
		log.Println(err)
	}
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			isCue, err := regexp.MatchString(`\.cue$`, event.Name)
			if err != nil {
				log.Println(err)
				continue
			}
			if (event.Op == fsnotify.Write || event.Op == fsnotify.Create) && isCue {
				log.Printf("%v modified", event.Name)
				err := generateImageAndSend(watcher.path, C)
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

func generateImageAndSend(p string, C chan diagram) error {
	cue := exec.Command("cue", "cmd", "genpuml")
	var outb, errb bytes.Buffer
	cue.Stdout = &outb
	cue.Dir = p
	cue.Stderr = &errb
	if err := cue.Run(); err != nil {
		C <- diagram{
			cue:    outb.Bytes(),
			cueErr: errb.Bytes(),
		}
		return fmt.Errorf("%v: %v", err, errb.String())
	}
	u, _ := url.Parse(config.PlantUMLServerAddress)
	u.Path = path.Clean(path.Join(u.Path, "/plantuml/svg/"))
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
	log.Printf("Senging new image on %v", C)
	puml, _ := format(bytes.NewReader(outb.Bytes()))
	C <- diagram{
		cue:      outb.Bytes(),
		cueErr:   errb.Bytes(),
		plantuml: []byte(puml),
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
