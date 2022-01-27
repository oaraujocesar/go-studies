package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/oaraujocesar/go-studies/blue-book/ch2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)

		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)

		convC, _ := tempconv.FahrenheitToCelsius(f)
		convF, _ := tempconv.CelsiusToFahrenheit(c)

		fmt.Printf("%s = %s, %s = %s\n", f, convC, c, convF)
	}
}
