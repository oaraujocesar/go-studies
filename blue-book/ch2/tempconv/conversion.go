package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CelsiusToFahrenheit(c Celsius) (f Fahrenheit, inputC Celsius) {
	return Fahrenheit(c*9/5 + 32), c
}

func FahrenheitToCelsius(f Fahrenheit) (c Celsius, inputF Fahrenheit) {
	return Celsius((f - 32) * 5 / 9), f
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
