package main

import (
	"fmt"
	"os"
	"strings"
)

// removes directory components and a .suffix.
// examples: a => a, a.go => a, a/b/c => c, a/b.c.go => b.c
func basename(s string) string {
	slash := strings.LastIndex(s, "/")

	s = s[slash+1:]

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
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
