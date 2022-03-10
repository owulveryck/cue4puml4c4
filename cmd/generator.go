package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"path"
)

func generateImageAndSend(p string, C chan diagram) error {
	cue := exec.Command("cue", "cmd", "genpuml")
	var outb, errb bytes.Buffer
	cue.Stdout = &outb
	cue.Dir = p
	cue.Stderr = &errb
	if err := cue.Run(); err != nil {
		C <- diagram{
			cue:      outb.Bytes(),
			cueErr:   errb.Bytes(),
			plantuml: []byte(`error`),
		}
		return fmt.Errorf("%v: %v", err, errb.String())
	}
	puml, _ := formatPlantuml(bytes.NewReader(outb.Bytes()))
	d := diagram{
		cueErr:   errb.Bytes(),
		plantuml: []byte(puml),
	}
	u, _ := url.Parse(config.PlantUMLServerAddress)
	u.Path = path.Clean(path.Join(u.Path, "/plantuml/svg/"))
	svgImage, err := callPlantumlServer(u, outb.Bytes())
	if err != nil {
		return err
	}
	d.svgImage = svgImage
	log.Printf("Senging new image on %v", C)
	C <- d
	return nil
}

func callPlantumlServer(plantumlURL *url.URL, payload []byte) ([]byte, error) {
	req := &http.Request{
		Method: http.MethodPost,
		URL:    plantumlURL,
		Header: map[string][]string{
			"Content-Type": {"text/plain"},
		},
		Body: io.NopCloser(bytes.NewReader(payload)),
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
