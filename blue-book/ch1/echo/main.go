// Echo shows your first argument on command line.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Running ", os.Args[0])

	for index, arg := range os.Args[1:] {
		fmt.Println(index, arg)
	}

	echo1()

	echo2()

	echo3()

	fmt.Printf("enlapsed %.6fs", time.Since(start).Seconds())
}

func echo1() {
	start := time.Now()

	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	secs := time.Since(start).Seconds()

	fmt.Printf("echo1 - %.7fs %s\n", secs, s)
}

func echo2() {
	start := time.Now()
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	secs := time.Since(start).Seconds()

	fmt.Printf("echo2 - %.7fs %s\n", secs, s)
}

func echo3() {
	start := time.Now()

	s := strings.Join(os.Args[1:], " ")

	secs := time.Since(start).Seconds()

	fmt.Printf("echo3 - %.7fs %s\n", secs, s)
}
