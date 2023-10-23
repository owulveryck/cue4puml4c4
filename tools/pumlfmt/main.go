package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	start := regexp.MustCompile(`{\s*$`)
	stop := regexp.MustCompile(`}\s*$`)

	scanner := bufio.NewScanner(os.Stdin)

	indent := 0
	for scanner.Scan() {
		curr := scanner.Text()
		if stop.MatchString(curr) {
			indent--
		}
		for i := 0; i < indent; i++ {
			fmt.Fprintf(os.Stdout, "\t")
		}
		if start.MatchString(curr) {
			indent++
		}
		fmt.Printf("%s\n", curr)
	}
}
