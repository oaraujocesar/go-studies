// Dup shows the text of every duplicated line
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// dup1()

	dup2()

	// dup3()
}

func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func dup2() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines("os.Stdin", os.Stdin, counts)
	} else {
		for _, arg := range files {
			file, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}

			countLines(arg, file, counts)

			defer file.Close()
		}
	}

	for line, filenames := range counts {
		fileCount := len(filenames)

		if fileCount == 1 {
			total := 0

			for _, count := range filenames {
				total += count
			}

			if total <= 1 {
				continue
			}
		}

		fmt.Printf("[Found in %d file(s)]\t%s\n", fileCount, line)
		for name, count := range filenames {
			fmt.Printf("\t%d hit(s) in %s\n", count, name)
		}
	}
}

func dup3() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(filename string, f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		// if the name is not there yet, this will add it.
		if counts[input.Text()] == nil {
			counts[input.Text()] = make(map[string]int)
		}

		// if it's already there, just increment the occurence
		counts[input.Text()][filename]++
	}
}
