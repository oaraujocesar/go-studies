package main

import (
	"fmt"
	"os"
)

// removes directory components and a .suffix.
// examples: a => a, a.go => a, a/b/c => c, a/b.c.go => b.c
func basename(s string) string {

	// Removes / and everything before it
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// Keeps everything before last .
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
		s := basename(arg)

		fmt.Printf("%s", s)
	}
}
