package main

import (
	"bufio"
	"fmt"
	"os"
)

// Line holds information about duplicate rows
type Line struct {
	Count int
	Files []string
}

func main() {
	counts := make(map[string]Line)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			defer f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts)
		}
	}
	for line, Line := range counts {
		if Line.Count > 1 {
			fmt.Printf("%d\t%s\t%v\n", Line.Count, line, Line.Files)
		}
	}
}

func countLines(f *os.File, counts map[string]Line) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		tmp := counts[input.Text()]
		tmp.Count++
		if !stringInSlice(f.Name(), tmp.Files) {
			tmp.Files = append(tmp.Files, f.Name())
		}
		counts[input.Text()] = tmp
	}
}

func stringInSlice(str string, slice []string) bool {
	for _, v := range slice {
		if str == v {
			return true
		}
	}
	return false
}
