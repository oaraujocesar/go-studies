package main

import (
	"fmt"
	"os"
)

// Inserts commas into a string that represents a non-negative decimal integer
func comma(s string) string {
	n := len(s)

	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	for _, arg := range os.Args[1:] {
		s := comma(arg)

		fmt.Println(arg)
		fmt.Printf("%s", s)
	}
}
