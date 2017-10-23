package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "stdin", counts)
	}
	for _, filename := range files {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f, filename, counts)
		f.Close()
	}
	for line, filenames := range counts {
		if len(filenames) > 1 {
			fmt.Printf("%d\t%s\n", len(filenames), line)
			fmt.Printf("Found in files: \n")
			for _, filename := range dedup(filenames) {
				fmt.Printf("\t%s\n", filename)
			}
		}
	}
}

func countLines(f *os.File, filename string, counts map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line] = append(counts[line], filename)
	}
}

func dedup(ar []string) []string {
	var last string
	var res []string

	for _, el := range ar {
		if el == last {
			continue
		}
		res = append(res, el)
		last = el
	}
	return res
}
