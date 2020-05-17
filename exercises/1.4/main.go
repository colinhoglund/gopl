package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type countsEntry struct {
	count int
	files map[string]bool
}

func newCountsEntry() *countsEntry {
	m := make(map[string]bool)
	return &countsEntry{files: m}
}

func main() {
	counts := make(map[string]*countsEntry)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts, arg)
			f.Close()
		}
	}

	for line, entry := range counts {
		if entry.count > 1 {
			var filesBuilder, sep string
			for f := range entry.files {
				filesBuilder += fmt.Sprintf("%s%s", sep, f)
				sep = ", "
			}

			fmt.Printf("line:\t%s\ncount:\t%d\nfiles:\t%v\n\n", line, entry.count, filesBuilder)
		}
	}
}

func countLines(f io.Reader, counts map[string]*countsEntry, filename string) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		text := input.Text()

		if counts[text] == nil {
			counts[text] = newCountsEntry()
		}

		counts[text].count++
		counts[text].files[filename] = true
	}
}
