// Echo shows your first argument on command line.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Running ", os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], " "))

	for index, arg := range os.Args[1:] {
		fmt.Println(index, arg)
	}
}
