package main

import (
	"bufio"
	"io"
	"strings"
)

func formatPlantuml(r io.Reader) (string, error) {
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
