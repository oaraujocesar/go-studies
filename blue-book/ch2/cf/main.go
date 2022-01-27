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
			fmt.Fprintf(os.Stderr, "cfk: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)

		convC, _ := tempconv.FahrenheitToCelsius(f)
		convF, _ := tempconv.CelsiusToFahrenheit(c)
		convK, _ := tempconv.CelsiusToKelvin(c)
		convFK, _ := tempconv.FahrenheitToKelvin(f)
		convKC, _ := tempconv.KelvinToCelsius(k)
		convKF, _ := tempconv.KelvinToFahrenheit(k)

		fmt.Printf("%s = %s, %s = %s\n", f, convC, c, convF)
		fmt.Printf("%s = %s, %s = %s, %s = %s, %s = %s\n", c, convK, f, convFK, k, convKC, k, convKF)
	}
}
