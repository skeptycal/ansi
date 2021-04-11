package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {
	fileName := "temp.go"

	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	fi, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fileName = fi.Name()

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	extractedLines := []string{}

	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "func ") || strings.HasPrefix(line, "type ") {
			extractedLines = append(extractedLines, line)
		}
	}

	fmt.Println(extractedLines)
}
